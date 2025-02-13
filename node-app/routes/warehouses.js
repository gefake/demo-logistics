const
    express = require('express'),
    rest_api = require('./_rest-api'),
    moment = require('moment'),
    mw = require('./_mw')

const router = new express.Router()
router.get('/', mw.Authorization, async(req, res) => {
    const products = await rest_api.get(`api/product`, {})
    rest_api.get('api/warehouse', {})
        .then(data => {
            console.log(data)
            res.render('admin/warehouse', {
                title: "Управление поставщиками",
                user: req.session.user || null,
                warehouses: data || [],
                products: products,
                banner: {"title": "Грузы и информация о них", "description": "Здесь находятся все текущие грузы"} })
        })
        .catch(error => {
            console.error('Ошибка причении данных:', error)
        });
})

router.post('/new', mw.Authorization, async(req, res) => {
    const user = req.session.user

    const body = {
        address: req.body.address,
        position: {lat: Number(req.body.address_lat), lon: Number(req.body.address_lon)}
    }

    console.log(body)

    rest_api.post('api/warehouse', body)
        .then(data => {
            console.log(data)
            res.redirect("/admin/warehouses")
        })
        .catch(error => {
            res.redirect("/admin/warehouses")
            console.error('Ошибка причении данных:', error)
        })
})

router.post('/delete/:id', mw.Authorization, async(req, res) => {
    rest_api.delete(`api/warehouse/${req.params.id}`, {})
        .then(data => {
            console.log(data)
            res.redirect("/admin/warehouses")
        })
        .catch(error => {
            res.redirect("/admin/warehouses")
            console.error('Ошибка причении данных:', error)
        })
})

module.exports = router