/**
 * 格式化字符串
 * @returns {String}
 */
String.prototype.format = function () {
    if (arguments.length === 0) return this;
    let param = arguments[0];
    let s = this;
    if (typeof (param) == 'object') {
        for (let key in param) s = s.replace(new RegExp("\\{" + key + "\\}", "g"), param[key]);
        return s;
    } else {
        for (let i = 0; i < arguments.length; i++) s = s.replace(new RegExp("\\{" + i + "\\}", "g"), arguments[i]);
        return s;
    }
}

/**
 * 超出长度字符串处理
 * @param maxLength
 * @param end
 * @returns {string|*}
 */
String.prototype.tooLong = function (maxLength, end = '…') {
    return this.length > maxLength ? this.substr(0, maxLength) + end : this;
}

/**
 * chunk数组
 * @param size
 * @returns {*[]|*}
 */
Array.prototype.chunk = function (size) {
    if (this.length <= 0 || size <= 0) {
        return this;
    }

    let chunks = [];

    for (let i = 0; i < this.length; i = i + size) {
        chunks.push(this.slice(i, i + size));
    }

    return chunks;
}

/**
 * 提取数组对象中的某个值
 * @param value
 * @param key
 * @returns {{}|*[]|*}
 */
Array.prototype.pluck = function (value, key) {
    if (this.length <= 0 || !value) {
        return this;
    }

    let newArr = [];
    let newDict = {};

    if (!key) {
        for (let i = 0; i < this.length; i++) {
            newArr.push(this[i][value]);
        }
        return newArr;
    } else {
        for (let i = 0; i < this.length; i++) {
            newDict[this[i][key]] = this[i][value];
        }
        return newDict;
    }
}

/**
 * 获取非重复数组
 * @returns {any[]}
 */
Array.prototype.unique = function () {
    return Array.from(new Set(this));
}

/**
 * 根据属性分组
 * @param dst
 * @returns {{}}
 */
Array.prototype.groupBy = function (dst) {
    let ret = {};

    this.map(value => {
        if (value.hasOwnProperty(dst)) {
            if (!ret.hasOwnProperty(value[dst])) {
                ret[value[dst]] = [];
            }
            ret[value[dst]].push(value);
        }
    });

    return ret;
};

/**
 * 数组合并字符串（去掉空值）
 * @param {string} limitStr
 * @returns {string}
 */
Array.prototype.joinWithNotEmpty = function (limitStr = ',') {
    let newArr = [];

    if (this.length > 0) {
        for (let i = 0; i < this.length; i++) {
            let item = this[i];
            if (!(item === '' || item === null || item === undefined)) {
                newArr.push(item);
            }
        }
    }

    return newArr.join(limitStr);
}

/**
 * 设置字典默认值
 * @param {{}} obj
 * @param {string} key
 * @param {*} value
 * @returns {{}}
 */
function setDefault(obj = {}, key = '', value = null) {
    if (!obj.hasOwnProperty(key)) {
        obj[key] = null;
        obj[key] = value;
    }
    return obj
}

/**
 * 生成随机颜色
 * @returns {string}
 */
const generateRandomHexColor = () => `#${Math.floor(Math.random() * 0xffffff).toString(16)}`

/**
 * 数组随机排序
 * @param arr
 */
Array.prototype.shuffle = () => this.sort(() => Math.random() - 0.5)

/**
 * 判断系统是否使用黑暗模式
 * @returns {boolean}
 */
const isDarkMode = () => window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches;

/**
 * 滚动到顶部
 * @param element
 * @returns {*}
 */
const scrollToTop = (element) => element.scrollIntoView({behavior: "smooth", block: "start"});

/**
 * 滚动到底部
 * @param element
 * @returns {*}
 */
const scrollToBottom = (element) => element.scrollIntoView({behavior: "smooth", block: "end"});

/**
 * 检测设备类型
 * @returns {string}
 */
const detectDeviceType = () => /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent) ? "Mobile" : "Desktop";

/**
 * 隐藏元素
 * @param el
 * @param removeFromFlow
 */
const hideElement = (el, removeFromFlow = false) => {
    removeFromFlow ? (el.style.display = 'none') : (el.style.visibility = 'hidden')
}

/**
 * 获取URL参数
 * @param key
 * @returns {string}
 */
const getParamByUrl = key => new URL(location.href).searchParams.get(key)

/**
 * 深拷贝
 * @param obj
 * @returns {any}
 */
const deepCopy = obj => JSON.parse(JSON.stringify(obj))

/**
 * 定时器
 * @param ms
 * @returns {Promise<unknown>}
 */
const wait = (ms) => new Promise((resolve) => setTimeout(resolve, ms))

/**
 * 异步定时器
 * @returns {Promise<void>}
 */
const asyncFn = async () => {
    await wait(1000)
    console.log('等待异步函数执行结束')
}

