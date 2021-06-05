package superspin

import (
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Block is a type for parsing text
type Block struct {
	id        string
	deepLevel int
	ref       string
	token     rune
	typ       string
}

// SuperSpin is the main type for spun text
type SuperSpin struct {
	openToken  rune
	closeToken rune
	orToken    rune
	text       string
	result     string
	seed       int64
	useSeed    bool
	blocks     []struct {
		id    string
		value []string
	}
}

// Seed use as default seed for SuperSpin
func (sp *SuperSpin) Seed(seed int64) {
	sp.seed = seed
}

func (sp *SuperSpin) parse() {
	tokens := []Block{}
	deepLevel := 0
	id := generateID()
	ids := [][]string{}
	ids = append(ids, []string{})
	ids[deepLevel] = append(ids[deepLevel], id)

	runes := []rune(sp.text)
	for i := 0; i < len(sp.text); i++ {
		char := runes[i]

		if char == sp.openToken {
			id := generateID()

			block := Block{
				id:  ids[deepLevel][len(ids[deepLevel])-1],
				ref: id,
				typ: "block",
			}

			deepLevel++
			ids = append(ids[:deepLevel], []string{})
			ids[deepLevel] = append(ids[deepLevel], id)
			tokens = append(tokens, block)
			continue
		}

		if char == sp.closeToken {
			deepLevel--
			continue
		}

		block := Block{
			id:        ids[deepLevel][len(ids[deepLevel])-1],
			token:     char,
			deepLevel: deepLevel,
			typ:       "char",
		}

		tokens = append(tokens, block)
	}

	blockGroup := map[string][]Block{}
	for i := 0; i < len(tokens); i++ {
		blockGroup[tokens[i].id] = append(blockGroup[tokens[i].id], tokens[i])
	}

	for key, val := range blockGroup {
		var block struct {
			id    string
			value []string
		}

		var s []string

		for i := 0; i < len(val); i++ {
			if val[i].typ == "char" {
				s = append(s, string(val[i].token))
			}

			if val[i].typ == "block" {
				s = append(s, "${"+val[i].ref+"}")
			}
		}

		block.id = key
		value := strings.Join(s, "")
		block.value = strings.Split(value, string(sp.orToken))
		sp.blocks = append(sp.blocks, block)
	}
}

var reBlock = regexp.MustCompile(`\$\{([\w\-\_]+)\}`)

// Spin create spin version of text
func (sp *SuperSpin) Spin(text string) string {
	sp.text = text
	sp.parse()
	// let seed = this._seed;
	// seed := sp.seed
	var spinnedBlocks []struct {
		id   string
		text string
	}

	for i := 0; i < len(sp.blocks); i++ {
		block := sp.blocks[i]
		id := block.id
		value := block.value

		randSeed, _ := strconv.ParseInt(id, 32, 64)
		if sp.seed > 0 {
			rand.Seed(sp.seed)
		} else {
			rand.Seed(int64(randSeed) + time.Now().UnixNano())
		}
		index := rand.Perm(len(value))
		text := value[index[0]]

		var spinnedBlock struct{ id, text string }
		spinnedBlock.id = id
		spinnedBlock.text = text
		spinnedBlocks = append(spinnedBlocks, spinnedBlock)
	}

	blockMap := make(map[string]string)
	for i := 0; i < len(spinnedBlocks); i++ {
		blockMap[spinnedBlocks[i].id] = spinnedBlocks[i].text
	}

	for {
		prevBlockLen := len(blockMap)
		for key, val := range blockMap {
			matched := reBlock.FindStringSubmatch(val)

			if len(matched) > 0 {
				rawID := matched[0]
				id := matched[1]
				blockMap[key] = strings.Replace(blockMap[key], rawID, blockMap[id], -1)
				delete(blockMap, id)
			}
		}

		if prevBlockLen == len(blockMap) {
			break
		}
	}

	lengthBlockMapVal := make(map[int]string)
	for key, val := range blockMap {
		lengthBlockMapVal[len(val)] = key
	}

	var keys []int
	for k := range lengthBlockMapVal {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	keyName := lengthBlockMapVal[keys[len(keys)-1]]
	sp.result = blockMap[keyName]

	return sp.result
}

func (sp *SuperSpin) String() (s string) {
	return sp.text
}

func generateID() string {
	rand.Seed(time.Now().UnixNano())
	numericID := rand.Int63()
	id := strconv.FormatInt(numericID, 32)
	return id
}

// New create new SuperSpin structure.
func New(params ...rune) SuperSpin {
	sp := SuperSpin{
		openToken:  '{',
		closeToken: '}',
		orToken:    '|',
	}

	if len(params) == 1 {
		sp.openToken = params[0]
	}

	if len(params) == 2 {
		sp.closeToken = params[1]
	}

	if len(params) == 3 {
		sp.orToken = params[2]
	}

	return sp
}
