echo "[ build start ]"
echo "mkdir dev"
mkdir dev
echo "mv dev/main dev/.main.old"
mv dev/main dev/.main.old
echo "cd src/main"
cd src/main
echo "go build"
go build
echo "mv main ../../dev"
mv main ../../dev
echo "[ build done ]"
