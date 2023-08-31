import request from './request'

function post(url, data) {
    return request({
        url: url,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: JSON.stringify(data)
    })
}

function get(url) {
    return request({
        url: url,
        method: 'GET'
    })
}

function buildQuery(form) {
    let query = ''
    for (let key in form) {
        query += key + '=' + form[key] + '&'
    }
    if (query == '') return query
    else return '?' + query.substring(0, query.length - 1)
}

export const config = {
    getAll: () => {
        return get('/config/all')
    },
    update: (key, value) => {
        return post('/config/update', {
            key, value
        })
    },
}

export const host = {
    add: (form) => {
        return post('/host/add', form)
    },
    edit: (id, form) => {
        return post('/host/' + id, form)
    },
    getList: (form) => {
        let query = buildQuery(form)
        return get('/host/list' + query)
    }
}

export const file = {
    getList: (id,path) => {
        return post('/file/list', {
            id,
            path
        })
    }
}