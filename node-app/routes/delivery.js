const
    express = require('express'),
    rest_api = require('./_rest-api'),
    moment = require('moment'),
    mw = require('./_mw')

const router = new express.Router()
router.get('/', mw.Authorization, async(req, res) => {
    const cargos = await rest_api.get('api/cargos?page=1&perPage=10', {})

    rest_api.get('api/delivery', {})
        .then(data => {
            // console.log(data)
            res.render('admin/delivery', {
                title: "Управление поставщиками",
                user: req.session.user || null,
                deliveries: data || [],
                cargos: cargos || [],
                banner: {"title": "Грузы и информация о них", "description": "Здесь находятся все текущие грузы"} })
        })
        .catch(error => {
            console.error('Ошибка причении данных:', error)
        });
})

router.post('/new', mw.Authorization, async(req, res) => {
    const user = req.session.user
    const cargo = await rest_api.get(`api/cargo/${req.body.cargo_id}`, {})

    const body = {
        driver_id: user.id,
        cargo_id: Number(req.body.cargo_id),
        arrival_date: moment().format(),
        departure_date: moment().format(),
        start_point: {lat: Number(req.body.address_lat), lon: Number(req.body.address_lon)},
        end_point: {lat: Number(req.body.address2_lat), lon: Number(req.body.address2_lon)},
        status: req.body.status
    }

    console.log(body)

    rest_api.post('api/delivery', body)
        .then(data => {
            // console.log(data)
            res.redirect("/admin/delivery")
        })
        .catch(error => {
            res.redirect("/admin/delivery")
            console.error('Ошибка причении данных:', error)
        })
})

router.post('/delete/:id', mw.Authorization, async(req, res) => {
    rest_api.delete(`api/delivery/${req.params.id}`, {})
        .then(data => {
            console.log(data)
            res.redirect("/admin/delivery")
        })
        .catch(error => {
            res.redirect("/admin/delivery")
            console.error('Ошибка причении данных:', error)
        })
})

module.exports = router