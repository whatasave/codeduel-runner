const prompt = require('readline-sync');
prompt.setDefaultOptions({
    
})

function readline() {
    return prompt.prompt()
}

module.exports = { readline };