package define

type Class struct {
    Cid       int64  `validate:"required||integer=10000,_"`
    Cname     string `validate:"required||string=1,5||unique"`
    BeginTime string `validate:"required||datetime=H:i"`
}

type Student struct {
    Uid          int64    `validate:"required||integer=10000,_"`
    Name         string   `validate:"required||string=1,5"`
    Age          int64    `validate:"required||integer=10,30"`
    Sex          string   `validate:"required||in=male,female"`
    Email        string   `validate:"email||user||vm"`
    PersonalPage string   `validate:"url"`
    Hobby        []string `validate:"array=_,2||unique||in=swimming,running,drawing"`
    CreateTime   string   `validate:"datetime"`
    Class        []Class  `validate:"array=1,3"`
}
