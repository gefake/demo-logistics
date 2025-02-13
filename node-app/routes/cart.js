const
    express = require('express'),
    rest_api = require('./_rest-api'),
    moment = require('moment'),
    mw = require('./_mw')

const router = new express.Router()

router.get('/', (req, res) => {
    let cartItems = JSON.parse(localStorage.getItem('cartItems')) || []
    res.render('cart', { title: "Корзина", user: req.session.user || null, items: cartItems })
})

router.post('/add', (req, res) => {
    let product = req.body.product
    product = JSON.parse(product)

    console.log(product)

    let cartItems = JSON.parse(localStorage.getItem('cartItems')) || []
    const existingItem = cartItems.find(item => item.product_id === product.id)

    if (existingItem) {
        existingItem.quantity = (existingItem.quantity || 0) + 1
    } else {
        cartItems.push({ product_id: product.id, name: product.name, supplier: product.supplier_id, quantity: 1, price: product.price })
    }

    localStorage.setItem('cartItems', JSON.stringify(cartItems))

    console.log(cartItems)

    // Отправка JSON-ответа для обновления корзины на клиенте
    res.json({ success: true, cartItems: cartItems })
})

router.post('/remove', (req, res) => {
    let product = req.body.product

    let cartItems = JSON.parse(localStorage.getItem('cartItems')) || []
    const index = cartItems.findIndex(item => item.name === product.name)

    if (index !== -1) {
        cartItems.splice(index, 1)
    }

    localStorage.setItem('cartItems', JSON.stringify(cartItems))
    res.json({ success: true, cartItems: cartItems })
})

router.post('/clear', (req, res) => {
    localStorage.setItem('cartItems', JSON.stringify([]))
    res.json({ success: true, cartItems: [] })
})

module.exports = router