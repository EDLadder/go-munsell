package organization

type Organization struct {
	UUID      string   `json:"uuid" bson:"_id,omitempty"`
	Name      string   `json:"name,omitempty" bson:"name,omitempty"`
	OwnerUUID string   `json:"owner_uuid" bson:"owner_uuid,omitempty"`
	Users     []string `json:"users,omitempty" bson:"users,omitempty"`
}

func NewOrganization(dto CreateOrganizationDTO) Organization {
	return Organization{
		Name:      dto.Name,
		OwnerUUID: dto.OwnerUUID,
	}
}

type CreateOrganizationDTO struct {
	Name      string `json:"name" bson:"name"`
	OwnerUUID string `json:"owner_uuid" bson:"owner_uuid"`
}
