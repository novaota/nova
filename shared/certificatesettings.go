package shared

type CertificateSettings struct {
	CACertificate     string
	CAKey             string
	DeviceCertificate string
	DeviceKey         string
}

//TODO: CHANGE BASEDIR TO DYNAMIC VALUE
//TODO: MAKE IT COMPATIBLE TO DOCKER CONTAINERS
var baseDir = "C:\\Users\\alf\\go\\src\\nova\\certificates\\"
var DefaultCertificateSettings CertificateSettings = CertificateSettings{
	CACertificate: baseDir + "penta.pem", //"ca-cert.pem",
	CAKey: baseDir + "penta.key", //"ca-key.pem",
	DeviceCertificate: baseDir + "device-cert.pem",
	DeviceKey: baseDir + "device-key.pem",
}