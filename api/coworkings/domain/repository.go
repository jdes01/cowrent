package domain

import r "api/utils/result"

type CoworkingRepository interface {
	SaveCoworking(*Coworking) r.Result[Coworking]
	GetCoworkings() r.Result[[]Coworking]
	GetCoworkingByUuid(*CoworkingUUID) r.Result[Coworking]
}
