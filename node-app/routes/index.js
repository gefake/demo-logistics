const
    express = require('express'),
    rest_api = require('./_rest-api'),
    _ = require('lodash'),
    tabulate = require('tabulate'),
    mw = require('./_mw'),
    axios = require('axios')
const moment = require('moment')

if (typeof localStorage === "undefined" || localStorage === null) {
    var LocalStorage = require('node-localstorage').LocalStorage
    localStorage = new LocalStorage('./scratch')
}

const router = new express.Router()

router.use('/user', require('./user'))
router.use('/order', require('./order'))
router.use('/cart', require('./cart'))

router.use('/admin', require('./admin'))
router.use('/admin/cargos', require('./cargos'))
router.use('/admin/delivery', require('./delivery'))
router.use('/admin/warehouses', require('./warehouses'))

router.get('/', async(req, res) => {
    const cats = await rest_api.get(`api/product/cats`, {})

    rest_api.get('api/supplier', {})
        .then(data => {
            // res.redirect('/admin/suppliers')
            // let cartItems = JSON.parse(localStorage.getItem('cartItems')) || []
            //
            // for (const supplier of data) {
            //     for (const obj of supplier.products) {
            //         const existingItem = cartItems.find(item => {
            //             return item.name === obj.name
            //         })
            //
            //         // data[obj].amount = cartItems[]
            //         if (existingItem) {
            //             existingItem.quantity = obj.quantity
            //         }
            //     }
            // }

            // console.log(data)

            res.render('main', {
                title: "Главная",
                user: req.session.user || null,
                suppliers: data,
                cats: cats
            })
        })
        .catch(error => {
            res.render('main', { title: "Главная", user: req.session.user || null })
            console.error('Ошибка причении данных:', error)
        })
})

module.exports = router
