package devicemgmtclient

import (
	"nova/devicemanagement/models"
	"nova/rest"
	"nova/shared"
)

func New(certSettings shared.CertificateSettings, endpoint string) *Client {
	return &Client{
		Client: rest.NewClient(certSettings, endpoint),
	}
}

type Client struct {
	*rest.Client
}

func (d *Client) Login(username string, password string) error {
	model := models.UserModel{
		Username: username,
		Password: password,
	}

	return d.Post(model, nil, "login", nil)
}

func (d *Client) AddDevice(model models.CarModel) error {
	return d.Post(model, nil, "devices")
}

func (d *Client) GetAllDevices() ([]models.CarModel, error) {
  result := &[]models.CarModel{}
	err := d.Get(result,"devices")

	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (d *Client) AddUpdate(model models.UpdateModel) error {
	return d.Post(model, nil,"updates")
}

func (d *Client) GetAllUpdates() ([]models.UpdateModel, error) {
	result := &[]models.UpdateModel{}
	err := d.Get(result,"updates")

	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (d *Client) AddOwner(model models.OwnerModel) error {
	return d.Post(model, nil,"owners")
}

func (d *Client) GetAllOwners() ([]models.OwnerModel, error) {
	result := &[]models.OwnerModel{}
	err := d.Get(result,"owners")

	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (d *Client) AddCarModel(model models.CarModelModel) error {
	return d.Post(model, nil,"carmodels")
}

func (d *Client) GetAllCarModels() ([]models.CarModelModel, error) {
	result := &[]models.CarModelModel{}
	err := d.Get(result,"carmodels")

	if err != nil {
		return nil, err
	}

	return *result, nil
}

func (d *Client) AddUpdateTask(model models.UpdateTaskModel) error {
	return d.Post(model, nil,"updates/tasks")
}

func (d *Client) GetAllUpdateTasks() ([]models.UpdateTaskModel, error){
	result := &[]models.UpdateTaskModel{}
	err := d.Get(result,"updates/tasks")

	if err != nil {
		return nil, err
	}

	return *result, nil
}