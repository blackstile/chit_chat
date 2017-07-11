

func sessionn(write http.ResponseWriter, request *http.Request)(sess data.Session, err error ){
    cookie, err := r.Cookie("_cookie")
    if err == nil {
        sess = data.Session{Uuid: cookie.value}
        if ok, _ := sess.Check(); !ok{
            err = errors.New("Invalid Session")
        }
    }
}