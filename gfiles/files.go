package gfiles

import (
	"bufio"
	"errors"
	"github.com/skypbc/goutils/gdirs"
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/gnum"
	"github.com/skypbc/goutils/internal"
	"io"
	"io/fs"
	"math"
	"os"
	pfilepath "path/filepath"
	"regexp"
	"strings"
	"sync"
)

func Exists(filepath string) bool {
	ok, _ := Exists2(filepath)
	return ok
}

func Exists2(filepath string) (bool, error) {
	if info, err := os.Stat(filepath); err == nil {
		if info.IsDir() {
			return false, nil
		}
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, gerrors.Wrap(err)
	}
}

func Size(filepath string) int64 {
	if size, err := Size2(filepath); err == nil {
		return size
	}
	return 0
}

func Size2(filepath string) (int64, error) {
	fd, err := os.Stat(filepath)
	if err != nil {
		return 0, gerrors.Wrap(err)
	}
	return fd.Size(), nil
}

func Read(filepath string) (data []byte, err error) {
	if data, err = os.ReadFile(filepath); err != nil {
		return nil, gerrors.Wrap(err)
	}
	return data, nil
}

func Read2(filepath string, size int64) (data []byte, fileSize int64, err error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, 0, gerrors.Wrap(err)
	}
	defer internal.Close(fd)

	stat, err := fd.Stat()
	if err != nil {
		return nil, 0, gerrors.Wrap(err)
	}

	fileSize = stat.Size()
	switch {
	case size == 0:
		return nil, fileSize, nil
	case size < 0:
		size = fileSize
	}

	data = make([]byte, gnum.Min(size, fileSize))

	if _, err = bufio.NewReader(fd).Read(data); err != nil && err != io.EOF {
		return nil, fileSize, gerrors.Wrap(err)
	}

	return data, fileSize, nil
}

func Read3(filepath string, out []byte) (fileSize int64, err error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return 0, gerrors.Wrap(err)
	}
	defer internal.Close(fd)

	stat, err := fd.Stat()
	if err != nil {
		return 0, gerrors.Wrap(err)
	}

	size := int64(len(out))

	fileSize = stat.Size()
	if size == 0 {
		return fileSize, nil
	} else if size < fileSize {
		return fileSize, gerrors.NewIncorrectParamsError().
			SetTemplate(`the output buffer is smaller than the file size: {bufSize} < {fileSize}`).
			AddInt64("bufSize", size).
			AddInt64("fileSize", fileSize)
	}

	if _, err = bufio.NewReader(fd).Read(out); err != nil && err != io.EOF {
		return fileSize, gerrors.Wrap(err)
	}

	return fileSize, nil
}

func Delete(filepath string) error {
	if exist, err := Exists2(filepath); err == nil && !exist {
		return nil
	}
	if err := os.Remove(filepath); err != nil {
		return gerrors.Wrap(err)
	}
	return nil
}

func Copy(toFile string, fromFile string, perm ...fs.FileMode) (int64, error) {
	if n, err := CopyWithBuffSize(toFile, fromFile, 8096, perm...); err != nil {
		return 0, gerrors.Wrap(err)
	} else {
		return n, nil
	}
}

func Move(toFile string, fromFile string) error {
	if err := os.Rename(fromFile, toFile); err != nil {
		return gerrors.Wrap(err)
	}
	return nil
}

func Create(filepath string, data []byte, perm ...fs.FileMode) (size int64, err error) {
	return Create2(filepath, data, 8192)
}

func Create2(filepath string, data []byte, bucket uint16, perm ...fs.FileMode) (size int64, err error) {
	fd, err := os.Create(filepath)
	if err != nil {
		return 0, gerrors.Wrap(err)
	}
	defer internal.Close(fd)

	dataSize := int64(len(data))
	if dataSize == 0 {
		return 0, nil
	}

	var start, end int64
	var count int

	for {
		end = start + int64(bucket)
		if dataSize < end {
			end = dataSize
		}

		if count, err = fd.Write(data[start:end]); err != nil && err != io.EOF {
			return size, gerrors.NewUnknownError(err)
		} else if count == 0 {
			break
		} else {
			start += int64(count)
		}
	}

	return dataSize, nil
}

func Create3(filepath string, source io.Reader, perm ...fs.FileMode) (size int64, err error) {
	fd, err := os.Create(filepath)
	if err != nil {
		return 0, gerrors.Wrap(err)
	}
	defer internal.Close(fd)
	return io.Copy(fd, source)
}

func CreateEmtpy(filepath string, size int64, perm ...fs.FileMode) (err error) {
	fd, err := os.Create(filepath)
	if err != nil {
		return gerrors.Wrap(err)
	}
	defer internal.Close(fd)

	if size == 0 {
		return nil
	}

	max, writed, bucket := int64(0), int64(0), int64(8*1024)
	data := make([]byte, bucket)

	for writed < size {
		if writed+bucket > size {
			max = size - writed
		} else {
			max = bucket
		}
		if n, err := fd.Write(data[:max]); err != nil && err != io.EOF {
			return gerrors.NewUnknownError(err)
		} else if n == 0 {
			break
		} else {
			writed += int64(n)
		}
	}

	return nil
}

func CopyWithBuffSize(toFile string, fromFile string, buffSize int64, perm ...fs.FileMode) (size int64, err error) {
	fromStat, err := os.Stat(fromFile)
	if err != nil {
		return 0, gerrors.Wrap(err)
	}
	if !fromStat.Mode().IsRegular() {
		return 0, gerrors.NewIncorrectParamsError().
			SetTemplate(`The "{from_file}" file isn't regular`).
			AddStr("from_file", fromFile)
	}

	if toStat, err := os.Stat(toFile); err == nil {
		if os.SameFile(fromStat, toStat) {
			// При копировании файла на самого себя, просто возвращаем размер файла
			return toStat.Size(), nil
		}
	}

	fromFd, err := os.Open(fromFile)
	if err != nil {
		return 0, gerrors.Wrap(err)
	}
	defer internal.Close(fromFd)

	toFd, err := os.OpenFile(toFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, fromStat.Mode())
	if err != nil {
		return 0, gerrors.Wrap(err)
	}
	defer internal.Close(toFd)

	buf := make([]byte, buffSize)
	for {
		n, err := fromFd.Read(buf)
		if err != nil && err != io.EOF {
			return 0, gerrors.Wrap(err)
		}
		if n == 0 {
			break
		}

		if _, err := toFd.Write(buf[:n]); err != nil {
			return 0, gerrors.Wrap(err)
		}
		size += int64(n)
	}

	if err = toFd.Sync(); err != nil {
		return size, gerrors.Wrap(err)
	}

	return size, nil
}

func Name(filepath string) string {
	fullname := FullName(filepath)
	if len(fullname) == 0 {
		return ""
	}
	return fullname[:len(fullname)-len(pfilepath.Ext(fullname))]
}

func Extension(filepath string) string {
	return pfilepath.Ext(FullName(filepath))
}

func FullName(filepath string) string {
	if filepath = pfilepath.Base(filepath); filepath == "." || filepath == "/" || filepath == "\\" {
		return ""
	}
	return filepath
}

func Dir(filepath string) string {
	if filepath = pfilepath.Dir(filepath); filepath == "." {
		return ""
	}
	return filepath
}

type ListOpts struct {
	Pattern     string
	Deep        *int
	TrimRoot    bool
	NormPathSep string
}

func List(root string, opts ...ListOpts) (res []string, err error) {
	var o ListOpts
	if len(opts) > 0 {
		o = opts[0]
	}
	if o.Deep == nil {
		deep := math.MaxInt
		o.Deep = &deep
	}
	var pattern *regexp.Regexp
	if len(o.Pattern) > 0 {
		if pattern, err = regexp.Compile(o.Pattern); err != nil {
			return nil, gerrors.Wrap(err)
		}
	}

	if err = gdirs.Traverse2(nil, root, 0, *o.Deep, func(entry fs.DirEntry, path string, deep int) error {
		if entry.IsDir() || (pattern != nil && !pattern.MatchString(path)) {
			return nil
		}
		if o.TrimRoot {
			tmp := []rune(root)
			if tmp[len(root)-1] == rune('/') || tmp[len(root)-1] == rune('\\') {
				tmp = []rune(path)[len(tmp):]
			} else {
				tmp = []rune(path)[len(tmp)+1:]
			}
			path = string(tmp)
		}
		if len(o.NormPathSep) > 0 {
			path = Normalize(path, o.NormPathSep)
		}
		res = append(res, path)
		return nil

	}); err != nil {
		return nil, gerrors.Wrap(err)
	}

	return res, nil
}

func Normalize(filepath string, sep ...string) string {
	sep_ := string(os.PathSeparator)
	if len(sep) > 0 && len(sep[0]) > 0 {
		sep_ = sep[0]
	}
	// Если filepath == "", то будет возвращена пустая строка, во всех остальных случаях, она будет
	// возвращена с заданным разделителем.
	filepath = strings.TrimSpace(filepath)
	if len(filepath) == 0 {
		return ""
	}
	if filepath = internal.CleanPath(filepath); len(filepath) == 0 {
		filepath = sep_
	}
	// Сохраняем корневой элемент, чтобы вернуть его в случае удаления, после чистки. Корневой элемент
	// есть только у е пустого пути.
	var root string
	if len(filepath) >= 2 && filepath[1] == ':' {
		if root = filepath[:2]; root == filepath {
			return root
		}
	} else if len(filepath) > 0 {
		root = sep_
		if len(filepath) == 1 && (filepath[0] == '/' || filepath[0] == '\\') {
			return root
		}
	}
	// Чистим путь от лишних символов
	filepath = pfilepath.Clean(filepath)
	if len(filepath) == 0 || filepath == "/" || filepath == "\\" || filepath == "." {
		return root
	}
	// Заменяем разделители на заданный
	if sep_ != "/" {
		filepath = strings.ReplaceAll(filepath, "/", sep_)
	}
	if sep_ != "\\" {
		filepath = strings.ReplaceAll(filepath, "\\", sep_)
	}
	// Если есть корень и путь не начинается с него, то добавляем его
	if len(root) > 0 && !strings.HasPrefix(filepath, root) {
		if root != sep_ {
			// windows
			filepath = root + sep_ + filepath
		} else {
			// unix
			filepath = root + filepath
		}
	}
	return filepath
}

func Segments(path string) (segments []string) {
	if len(path) == 0 || path == "/" || path == "\\" {
		segments = append(segments, "")
	} else {
		segments = strings.Split(Normalize(path, "/"), "/")
	}
	return segments
}

var (
	exePathOnce sync.Once
	exePathVal  string
	exePathErr  error
)

// Executable возвращает абсолютный путь к директории, где лежит исполняемый файл. Вычисляется один раз и кешируется.
func Executable() (string, error) {
	exePathOnce.Do(func() {
		exePathVal, exePathErr = os.Executable()
		if exePathErr != nil {
			return
		}
		// Нормализуем путь: раскрываем симлинки, чтобы получить "реальное" расположение бинарника
		exePathVal, exePathErr = pfilepath.EvalSymlinks(exePathVal)
	})
	return exePathVal, exePathErr
}
