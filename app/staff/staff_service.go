package staff

type StaffRepository interface {
	PostStaff(req Staff) error
}

func NewApp(staffRepository StaffRepository) StaffService {
	return StaffService{
		staffRepo: staffRepository,
	}
}

type StaffService struct {
	staffRepo StaffRepository
}
