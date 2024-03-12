package merkledag

import "hash"

func Add(store KVStore, node Node, h hash.Hash) []byte {

	if node.Type() == FILE {

		fileNode := node.(File)
		value := fileNode.Bytes()
		err := store.Put(h.Sum(nil), value)
		if err != nil {
			return nil
		}
		return h.Sum(nil)
	} else if node.Type() == DIR {

		dirNode := node.(Dir)
		it := dirNode.It()
		var childHashes [][]byte
		for it.Next() {
			childNode := it.Node()
			childHash := Add(store, childNode, h)
			if childHash == nil {

				return nil
			}
			childHashes = append(childHashes, childHash)
		}

		for _, hash := range childHashes {
			h.Write(hash)
		}
		return h.Sum(nil)
	}

	return nil
}
