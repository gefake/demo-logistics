const
    express = require('express'),
    rest_api = require('./_rest-api'),
    moment = require('moment'),
    mw = require('./_mw')

const router = new express.Router()
router.post('/products', mw.Authorization, (req, res) => {
    let cartItems = JSON.parse(localStorage.getItem('cartItems')) || []
    let user = req.session.user

    // for (const obj of cartItems) {
    //     console.log(obj)
    // }

    const body = {
        address: "Test",
        client_id: Number(user.id),
        order_date: "2006-01-02T15:04:05Z",
        status: "В процессе",
        orderItems: cartItems
    }

    console.log(body)

    rest_api.post('api/order', body)
        .then(data => {
            // res.redirect('/admin/suppliers')
            localStorage.setItem('cartItems', JSON.stringify([]))
            res.render('message', { text: 'Спасибо! Ваш заказ успешно был принят в работу', type: 'success', btext: "На главную", back: '/'})
        })
        .catch(error => {
            // res.redirect('/admin/suppliers')
            console.error('Ошибка причении данных:', error);
        });
})

router.post('/delete/:id', mw.Authorization, async(req, res) => {
    rest_api.delete(`api/order/${req.params.id}`, {})
        .then(data => {
            console.log(data)
            res.redirect("/admin")
        })
        .catch(error => {
            res.redirect("/admin")
            console.error('Ошибка причении данных:', error)
        })
})

module.exports = router