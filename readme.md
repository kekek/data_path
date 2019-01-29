### 获取文件路径的方法

### 利用 runtime 获取当前文件路径， 获取测试数据的方法

```
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

```