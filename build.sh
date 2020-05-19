echo "[ build start ]"
echo "mv main .main.old"
mv main .main.old
echo "cd src/main"
cd src/main
echo "go build"
go build
echo "mv main ../../"
mv main ../../
echo "[ build done ]"
