const shortid = require('shortid');
const _ = require('lodash');
const gen = require('random-seed');

class Superspin {
  constructor(openToken = '{', closingToken = '}', orToken = '|') {
    this.openToken = openToken;
    this.closingToken = closingToken;
    this.orToken = orToken;
    this.blocks = {};
    this._seed = shortid.generate();
  }

  seed(id) {
    this.seed = id;
  }

  _normalizeBlock(tokens, key) {
    let ret = {};
    ret.id = key;
    ret.value = _.chain(tokens)
      .map(data => data.type === 'char' ? data.token : `$\{${data.ref}\}`)
      .join('')
      .split(this.orToken)
      .value();
    return ret;
  }

  parse(s) {
    let arr = s.split('');
    let tokens = [];
    let deepLevel = 0;
    let id = shortid.generate();
    let ids = [];
    ids[0] = [];
    ids[deepLevel].push(id);

    for(let i in arr) {
      let char = arr[i];
      if (char === this.openToken) {
        let id = shortid.generate();

        let block = {
          id: _.last(ids[deepLevel]),
          ref: id,
          type: 'block'
        };

        deepLevel++;
        ids[deepLevel] = [];
        ids[deepLevel].push(id);
        tokens.push(block);
        continue;
      }

      if (char === this.closingToken) {
        deepLevel--;
        continue;
      }

      let block = {
        id: _.last(ids[deepLevel]),
        token: char,
        deepLevel,
        type: 'char'
      };

      tokens.push(block);
    }

    let blocks = _.chain(tokens)
      .groupBy('id')
      .map(this._normalizeBlock.bind(this))
      .value();

    this.blocks = blocks;
  }

  blocksLength() {
    return this.blocks.length;
  }

  _getBlock(id) {
    return this.blocks[id];
  }

  spin(s, withSeed = false) {
    this.parse(s);
    let seed = this._seed;
    let spinnedBlock = [];
    for (let x in this.blocks) {
      let block = this.blocks[x];
      let {id, value} = block;
      let rand = withSeed ? gen(seed) : gen(id);
      let valueLength = value.length;
      let index = rand.range(valueLength);
      let text = value[index];
      spinnedBlock.push({id, text});
    }

    let blockPattern = new RegExp('\\$\\{([\\w\\-\\_]+)\\}', 'i');
    let targetText = spinnedBlock[0].text;
    var result;
    while((result = blockPattern.exec(targetText)) !== null) {
      let id = result[1];
      let findBlock = _.find(spinnedBlock, {id});
      let pattern = new RegExp('\\$\\{'+id+'\\}', 'ig');
      targetText = _.replace(targetText, pattern, findBlock.text);
    }

    return targetText;
  }
}

module.exports = Superspin;
