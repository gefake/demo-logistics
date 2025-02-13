const
    express = require('express'),
    rest_api = require('./_rest-api'),
    mw = require('./_mw')

const router = new express.Router()

router.post('/login', async(req, res) => {
    var user = req.body.username
    var password = req.body.password

    const body = {
        login: user,
        password: password
    }

    rest_api.post('auth/login', body)
        .then(data => {
            req.session.user = data
            res.redirect('/')

            console.log(req.session.user)
        })
        .catch(error => {
            res.redirect('/user/login')
            console.error('Ошибка причении данных:', error);
        });
})

router.post('/reg', async(req, res) => {
    const body = {
        username: req.body.username,
        firstname: req.body.name,
        lastname: req.body.lastname,
        email: req.body.email,
        phone: req.body.phone,
        password: req.body.password
    }

    rest_api.post('auth/register', body)
        .then(data => {
            // req.session.user = body
            res.redirect('/')
            console.log(data)
        })
        .catch(error => {
            res.redirect('/user/reg')
            console.error('Ошибка причении данных:', error);
        });
})

router.get('/login', (req, res) => {
    res.render('user/login', {title: "Вход в систему"})
})

router.get('/reg', (req, res) => {
    res.render('user/reg', {title: "Регистрация"})
})

module.exports = router