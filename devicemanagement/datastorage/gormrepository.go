package datastorage

import (
	"github.com/jinzhu/gorm"
	"nova/devicemanagement/datastorage/entities"

)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(dialect string, connectionString string) *GormRepository {
	var repo = &GormRepository{}

	repo.initDb(dialect, connectionString)

	return repo
}

func (repo *GormRepository) initDb(dialect string, connectionString string) {
	db, error := gorm.Open(dialect, connectionString)

	if error != nil {
		db.Close()
		panic("Could not connect to database")
	}

	repo.db = db
}

func (repo *GormRepository) GetAll(out interface{}) {
	repo.db.Find(out)
}

func (repo *GormRepository) Get(id int, out interface{}) {
	repo.db.First(&out, id)
}

func (repo *GormRepository) AddOrUpdate(model entities.Entity) error {
	err := model.Validate()
	if err != nil {
		return err
	}

	repo.db.FirstOrCreate(&model, model.GetID())
	return nil
}

func (repo *GormRepository) First(out interface{}, predicate interface{}) {
	repo.db.First(out, predicate)
}

func (repo *GormRepository) GetWhere(out interface{}, predicate interface{}) {
	repo.db.Where(predicate).Find(out)
}

func (repo *GormRepository) Add(model interface{}) error {
	validatableModel, ok := model.(entities.Entity)

	if ok {
		err := validatableModel.Validate()
		if err != nil {
			return err
		}
	}

	repo.db.Create(model)
	return nil
}

func (repo *GormRepository) Update(model entities.Entity) error {
	err := model.Validate()
	if err != nil {
		return err
	}

	repo.db.UpdateColumns(&model)
	return nil
}

func (repo *GormRepository) Delete(model entities.Entity) {
	repo.db.Delete(model)
}

func (repo *GormRepository) Close() {
	repo.db.Close()
}

func (repo *GormRepository) CreateSchema() {
	// create tablesentities.
	repo.db.AutoMigrate(&entities.Device{}, &entities.DeviceModel{}, &entities.Owner{}, &entities.Update{}, &entities.UpdateTask{}, &entities.User{})

	// foreign keys are not supported by auto migrate
	// so lets do it by our own
	repo.db.Model(&entities.Device{}).AddForeignKey("device_model_id", "device_models(id)", "NO ACTION", "NO ACTION")
	repo.db.Model(&entities.Update{}).AddForeignKey("suitable_device_id", "device_models(id)", "NO ACTION", "NO ACTION")
	repo.db.Model(&entities.UpdateTask{}).AddForeignKey("device_id", "devices(id)", "NO ACTION", "NO ACTION")
	repo.db.Model(&entities.UpdateTask{}).AddForeignKey("update_id", "updates(id)", "NO ACTION", "NO ACTION")

	//constraints
	repo.db.Model(&entities.User{}).AddUniqueIndex("idx_unique_username", "username")
}

func (repo *GormRepository) IsSchemaCreated() bool {
	return repo.db.HasTable(&entities.Device{})
}
