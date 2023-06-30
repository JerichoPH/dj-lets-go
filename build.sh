echo "开始打包……" &&
echo "初始化打包输出目录" &&
rm -rf "./out/$2" &&
mkdir -p "./out/$2" &&
echo "编译【程序：dj-lets-go】【平台：$1】【版本：$2】" &&
CGO_ENABLED=0 GOOS="$1" GOARCH=amd64 go build -a -o "./out/$2/" "dj-lets-go" &&
echo "打包：模板文件" &&
cp -r ./templates "./out/$2/" &&
echo "打包配置文件" &&
rm -rf "./out/$2/settings" &&
mkdir "./out/$2/settings" &&
cp ./settings/app.ini.exa "./out/$2/settings/" &&
cp ./settings/db.ini.exa "./out/$2/settings/" &&
echo "打包静态文件" &&
cp -r ./static "./out/$2" &&
echo "编译完成"