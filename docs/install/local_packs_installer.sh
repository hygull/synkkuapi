#Whenever we make any changes to the files inside our local packages, then we need to build and install it

echo ""
echo "********************       Synkku       ***********************"
echo "----------------- INSTALLING LOCAL PACKAGES -------------------"
echo ""

cd $GOPATH/src/synkkuapi/conf
go build
go install 
go fmt
echo "1) conf successfully installed"

cd $GOPATH/src/synkkuapi/views 
go build
go install 
go fmt
echo "2) views successfully installed"

cd $GOPATH/src/synkkuapi/controllers
go build
go install 
go fmt
echo "3) controllers successfully installed"

# cd $GOPATH/src/synkkuapi/controllers/validation
# go build
# go install 
# go fmt
# echo "3) controllers/validation successfully installed"

cd $GOPATH/src/synkkuapi/routers/
go build
go install 
go fmt
echo "4) routers successfully installed"

cd $GOPATH/src/synkkuapi/
go build
go fmt main.go
echo "5) synkku successfully built (synkku.exe modified)"
echo ""
