package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	r "math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type outfmt struct {
	Strings []string
	Count   int
	Length  int
}

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/api", APIHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), s)
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	var count, length int
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	cnt := r.URL.Query().Get("count")
	ln := r.URL.Query().Get("len")
	if ln == "" {
		length = 10
	} else {
		stat := TypeConv(ln, w, &length)
		if stat == false {
			return
		}
	}
	if cnt == "" {
		count = 1
	} else {
		stat := TypeConv(cnt, w, &count)
		if stat == false {
			return
		}
	}
	lnt := length / 2
	if lnt >= 100 {
		w.WriteHeader(400)
		w.Write([]byte("length should be less than 200"))
		return
	}
	if count >= 200 {
		w.WriteHeader(400)
		w.Write([]byte("count should be less than 200"))
		return
	}
	input := MakeRand(lnt, count, length)
	out, err := GenerateJSON(input)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(out))
}

func GenerateJSON(j outfmt) (out string, err error) {
	var o []byte
	o, err = json.Marshal(j)
	out = string(o)
	return out, err
}

func MakeRand(lnt, count, length int) (out outfmt) {
	c := make([]byte, lnt)
	for i := 1; i <= count; i++ {
		rand.Read(c)
		rndout := hex.EncodeToString(c)
		if len(rndout) != length*2 {
			r.Seed(time.Now().UnixNano())
			rndout += strconv.Itoa(r.Intn(9-0+1) + 0)
		}
		out.Strings = append(out.Strings, rndout)
		out.Count = count
		out.Length = length
	}
	return out
}

func TypeConv(data string, w http.ResponseWriter, out *int) (stat bool) {
	stat = true
	o, err := strconv.Atoi(data)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data Type"))
		stat = false
	}
	*out = o
	return stat
}
