const
    express = require('express'),
    rest_api = require('./_rest-api'),
    mw = require('./_mw')

const router = new express.Router()
router.get('/', mw.Authorization, (req, res) => {
    rest_api.get(`api/order`, {})
        .then(data => {
            // res.redirect('/admin/suppliers')
            res.render('admin/panel', { orders: data, title: "Управление заказами", user: req.session.user || null, banner: {"title": "Панель управления заказами", "description": "Здесь можно просматривать все текущие заказы и их подробную карту с точками"} })
            console.log(data)

            // for (const obj of data) {
            //     console.log(obj.OrderItems)
            // }

        })
        .catch(error => {
            console.error('Ошибка причении данных:', error);
        });
})

router.get('/order', mw.Authorization, (req, res) => {
    res.render('admin/order', { title: "Управление заказами", user: req.session.user || null, banner: {"title": "Подробная информация о доставке", "description": "Здесь можно просматривать всю информацию о конкретном заказе и конкретные маршруты следования из точки А в точку Б"} })
})

router.get('/suppliers', mw.Authorization, (req, res) => {

    rest_api.get('api/supplier', {})
        .then(data => {
            // res.redirect('/admin/suppliers')
            console.log(data)

            res.render('admin/suppliers', {
                title: "Управление поставщиками",
                user: req.session.user || null,
                suppliers: data,
                banner: {"title": "Поставщики и информация о них", "description": "Здесь находятся все доступные вам поставщики. Также, при необходимости можно добавить новых"} })
        })
        .catch(error => {
            res.redirect('/admin/suppliers')
            console.error('Ошибка причении данных:', error);
        });
})

router.get('/supplier/:id/products', mw.Authorization, async(req, res) => {
    const id = req.params.id
    const warehouses = await rest_api.get(`api/warehouse`, {})

    rest_api.get(`api/supplier/${id}`, {})
        .then(data => {
            // res.redirect('/admin/suppliers')
            console.log(data)

            res.render('admin/products', {
                title: "Управление продуктами поставщика",
                user: req.session.user || null,
                suppliers: data,
                warehouses: warehouses,
                banner: {"title": "Управление продуктами поставщика", "description": "Здесь находятся все доступные продукты для поставщика. Также, при необходимости можно добавить новые"} })
        })
        .catch(error => {
            res.redirect('/admin/suppliers')
            console.error('Ошибка причении данных:', error);
        });
})

router.post('/suppliers-add', async(req, res) => {
    const body = {
        name: req.body.name,
        contact: req.body.contact,
        email: req.body.email,
    }

    console.log(body)

    rest_api.post('api/supplier', body)
        .then(data => {
            res.redirect('/admin/suppliers')
            console.log(data)
        })
        .catch(error => {
            res.redirect('/admin/suppliers')
            console.error('Ошибка причении данных:', error);
        });
})

router.post('/suppliers-remove/:id', mw.Authorization, async(req, res) => {
    rest_api.delete(`api/supplier/${req.params.id}`, {})
        .then(data => {
            console.log(data)
            res.redirect("/admin/suppliers")
        })
        .catch(error => {
            res.redirect("/admin/suppliers")
            console.error('Ошибка причении данных:', error)
        })
})

// products

router.post('/supplier/:id/product-add', async(req, res) => {
    const id = req.params.id

    const body = {
        name: req.body.name,
        category: req.body.category,
        unit: req.body.unit,
        price: Number(req.body.price),
        quantity: Number(req.body.quantity),
        warehouse_id: Number(req.body.warehouse_id),
        description: req.body.description,
        supplier_id: Number(id)
    }

    console.log(body)

    rest_api.post(`api/product`, body)
        .then(data => {
            res.redirect(`/admin/supplier/${id}/products`)
            console.log(data)
        })
        .catch(error => {
            // res.redirect(`/admin/supplier/${id}/products`)
            console.error('Ошибка причении данных:', error);
        });
})

router.post('/supplier/:id/product-remove/:product_id', async(req, res) => {
    const id = req.params.id
    const product_id = req.params.product_id

    const body = {
        id: Number(id)
    }

    rest_api.delete(`api/product/${Number(product_id)}`, body)
        .then(data => {
            res.redirect(`/admin/supplier/${id}/products`)
        })
        .catch(error => {
            // res.redirect(`/admin/supplier/${id}/products`)
            console.error('Ошибка причении данных:', error);
        });
})

module.exports = router