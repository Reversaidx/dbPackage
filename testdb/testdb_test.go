package testdb

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDatabase(t *testing.T) {
	t.Run("Complex", func(t *testing.T) {
		var test Database
		err := test.New()
		require.Nil(t, err)
		err = test.Put([]byte("aaa"), []byte("100"))
		require.Nil(t, err)
		err = test.Put([]byte("bbb"), []byte("200"))
		require.Nil(t, err)
		err = test.Put([]byte("ccc"), []byte("300"))
		require.Nil(t, err)
		err = test.Put([]byte("ddd"), []byte("400"))
		require.Nil(t, err)

		//test put
		val, ok := test.Get([]byte("aaa"))
		require.Nil(t, ok)
		require.Equal(t, []byte("100"), val)
		// test delete

		err = test.Delete([]byte("aaa"))
		require.Nil(t, err)
		_, ok = test.Get([]byte("aaa"))
		// should be not found
		require.NotNil(t, ok)

		//update
		err = test.Put([]byte("bbb"), []byte("250"))
		require.Nil(t, err)
		val, ok = test.Get([]byte("bbb"))
		require.Nil(t, ok)
		require.Equal(t, []byte("250"), val)

		//flush
		test.Flush()
		_, ok = test.Get([]byte("bbb"))
		// should be not found
		require.NotNil(t, ok)

		//testStat

		stat := test.Stats()

		require.Equal(t, "inserts:4\nupdates:1\ndeletes:1\nmiss:2\nhit:2\n", stat)

	})

}
