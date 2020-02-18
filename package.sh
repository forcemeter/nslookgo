# compile for version
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

# cross_compiles
make -f ./Makefile.cross-compiles

pkg=./packages
app=nslookgo
ver=0.0.1

os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle'

rm -rf $pkg
mkdir  $pkg

for os in $os_all; do
    for arch in $arch_all; do
        name="bin/${app}_${os}_${arch}"
        zip="${name}_${ver}.zip"

        #Windows file extension
        if [ "x${os}" = x"windows" ]; then
            name="${name}.exe"
        fi  

        if [ ! -f "./${name}" ]; then
            continue
        fi

        zip -rq ${zip} ${name}
        mv ${zip} ${pkg}

        echo "${pkg}/${zip} done"
    done
done
