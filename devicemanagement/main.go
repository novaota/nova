package main

func main() {
	api := NewDeviceManagementApi( 8080)
	api.SetName("Device Management Api")
	api.SetupPostgresDatabase("host=localhost port=5432 user=postgres dbname=devicemanagement password=passwort sslmode=disable")
	//	dm.UseTls("C:/Users/alf/go/src/updateservice/certificates/penta.crt", "C:/Users/alf/go/src/updateservice/certificates/penta.key")
	api.StartServing()
}
