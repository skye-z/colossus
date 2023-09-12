const options = (group, addAll) => {
    let retMap = {}
    let dataMap = {
        platform: [
            {
                label: "Linux",
                value: "Linux"
            },
            {
                label: "MacOS",
                value: "MacOS"
            },
            {
                label: "Windows",
                value: "Windows"
            },
        ],
        system: [
            {
                label: "RedHat",
                value: "RedHat"
            },
            {
                label: "Ubuntu",
                value: "Ubuntu"
            },
            {
                label: "CentOS",
                value: "CentOS"
            },
            {
                label: "Debian",
                value: "Debian"
            },
            {
                label: "Oracle",
                value: "Oracle"
            },
            {
                label: "SUSE",
                value: "SUSE"
            },
            {
                label: "Fedora",
                value: "Fedora"
            },
            {
                label: "Arch",
                value: "Arch"
            },
            {
                label: "FreeBSD",
                value: "FreeBSD"
            }
        ],
        region: [
            {
                label: "内网",
                value: "内网"
            },
            {
                label: "亚洲",
                value: "亚洲"
            },
            {
                label: "欧洲",
                value: "欧洲"
            },
            {
                label: "北美洲",
                value: "北美洲"
            },
            {
                label: "南美洲",
                value: "南美洲"
            },
            {
                label: "大洋洲",
                value: "大洋洲"
            },
            {
                label: "非洲",
                value: "非洲"
            },
        ],
        usage: [
            {
                label: "个人学习",
                value: "个人学习"
            },
            {
                label: "个人项目",
                value: "个人项目"
            },
            {
                label: "个人测试",
                value: "个人测试"
            },
            {
                label: "家庭服务",
                value: "家庭服务"
            },
            {
                label: "企业项目",
                value: "企业项目"
            },
            {
                label: "企业测试",
                value: "企业测试"
            },
        ],
        period: [
            {
                label: "长期有效",
                value: "长期有效"
            },
            {
                label: "近7天过期",
                value: "近7天过期"
            },
            {
                label: "近15天过期",
                value: "近15天过期"
            },
            {
                label: "近30天过期",
                value: "近30天过期"
            },
        ]
    }
    for (let i in group) {
        let key = group[i];
        if (addAll != undefined && addAll == true) dataMap[key].unshift({
            label: "All",
            value: ""
        })
        retMap[key] = dataMap[key]
    }
    return retMap
}

export default options