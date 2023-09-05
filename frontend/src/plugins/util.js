const util = {
    // 构建文件路径
    buildFilePath(oldList, newList, index) {
        let path = ''
        if (newList.length > 6) {
            for (let i = 0; i <= newList.length - 5; i++) {
                path += '/' + newList[i]
            }
        }
        for (let i = (newList.length > 6 ? 2 : 0); i <= index; i++) {
            path += '/' + oldList[i]
        }
        if (path.length > 1) {
            let check = path.substring(0, 2);
            if (check == '/~' || check == '//') path = path.substring(1)
        }
        return path
    },
    // 解析文件数据
    analysisFile(data) {
        if (data.name == '.' || data.name == '..') return null;
        if (data.type == 1) data.typeName = '文件'
        else if (data.type == 2) data.typeName = '目录'
        else if (data.type == 3) data.typeName = '程序'
        else if (data.type == 4) data.typeName = '链接'

        if (data.name.indexOf('->') != -1) {
            data.type = 4
            data.typeName = '链接'
            data.path = data.name.substring(data.name.indexOf('->') + 2, data.name.length)
            data.name = data.name.substring(0, data.name.indexOf('->'))
        }

        data.owner = data.user + ' / ' + data.group
        if (data.type == 2 || data.type == 4) data.sizeText = '-'
        else data.sizeText = util.bytes2Size(data.size)
        return data
    },
    // 字节转文件大小
    bytes2Size(number) {
        const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
        if (number === 0) return '0 Byte';
        const i = parseInt(Math.floor(Math.log(number) / Math.log(1024)));
        return Math.round(number / Math.pow(1024, i), 2) + ' ' + sizes[i];
    },
    // API函数错误
    funcError(err, loading, action) {
        loading = false
        console.log(err)
        window.$message.warning(action + '失败, 发生意料之外的错误')
    },
    funcErrorNoLoading(err, action) {
        console.log(err)
        window.$message.warning(action + '失败, 发生意料之外的错误')
    },
    funcErrorDIYTips(err, tips, action) {
        console.log(err)
        tips.type = 'warning'
        tips.content = action + "失败, 发生意料之外的错误"
        setTimeout(() => {
            tips.destroy()
        }, 2000)
    }
}

export default util