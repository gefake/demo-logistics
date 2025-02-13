const
    express = require('express'),
    rest_api = require('./_rest-api'),
    mw = require('./_mw')

const router = new express.Router()
router.get('/', mw.Authorization, async(req, res) => {
    const orders = await rest_api.get('api/order', {})

    rest_api.get('api/cargos?page=1&perPage=10', {
    })
        .then(data => {
            console.log(data)
            res.render('admin/cargos', {
                title: "Управление поставщиками",
                user: req.session.user || null,
                cargos: data.cargos,
                orders: orders || [],
                banner: {"title": "Грузы и информация о них", "description": "Здесь находятся все текущие грузы"} })
        })
        .catch(error => {
            res.redirect("/admin/cargos")
            console.error('Ошибка причении данных:', error)
        })
})

router.post('/new', mw.Authorization, (req, res) => {
    const user = req.session.user

    const body = {
        client_id: user.id,
        order_id: Number(req.body.order_id),
        name: req.body.name,
        description: req.body.description,
        status: "В обработке",
        weight: Number(req.body.weight)
    }

    rest_api.post('api/cargo', body)
        .then(data => {
            console.log(data)
            res.redirect("/admin/cargos")
        })
        .catch(error => {
            res.redirect("/admin/cargos")
            console.error('Ошибка причении данных:', error)
        })
})

router.post('/delete/:id', mw.Authorization, async(req, res) => {
    rest_api.delete(`api/cargo/${req.params.id}`, {})
        .then(data => {
            console.log(data)
            res.redirect("/admin/cargos")
        })
        .catch(error => {
            res.redirect("/admin/cargos")
            console.error('Ошибка причении данных:', error)
        })
})

module.exports = router