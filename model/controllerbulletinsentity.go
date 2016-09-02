package model

type ControllerBulletinsEntity struct {
	Bulletins                  []BulletinEntity `json:"bulletins"`                  // System level bulletins to be reported to the user.
	ControllerServiceBulletins []BulletinEntity `json:"controllerServiceBulletins"` // Controller service bulletins to be reported to the user.
	ReportingTaskBulletins     []BulletinEntity `json:"reportingTaskBulletins"`     // Reporting task bulletins to be reported to the user.
}
