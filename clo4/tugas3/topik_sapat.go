package main

import "fmt"

type tag struct {
	topik  string
	banyak int
}

type tabTopic = [100]string

type tabTag = [100]tag

func main() {
	var topic tabTopic
	var tag tabTag
	isi(&topic, &tag)
	trending(tag)
}

func cari(topic *tabTopic, kata string) int {
	var i, k int
	for i < len(topic) {
		if kata == topic[i] {
			k++
		}
		i++
	}
	return k
}

func isi(topic *tabTopic, tag *tabTag) {
	var i, j, k, l int
	var kata string
	var n string
	for i < len(topic) && n != "." {
		fmt.Scan(&n)
		if n != "." {
			topic[i] = n
		}
		i++
	}
	for j < len(topic) {
		if topic[j] != "" {
			fmt.Println(topic)
			kata = topic[j]
			tag[l].topik = kata
			tag[l].banyak = cari(topic, kata)
			l++
			for k < len(topic) {
				if topic[k] == kata {
					topic[k] = ""
				}
				k++
			}
		}
		j++
		k = 0
	}
}

func trending(tag tabTag) {
	var max, i, id int
	for i < len(tag) {
		if max < tag[i].banyak {
			max = tag[i].banyak
			id = i
		}
		i++
	}
	fmt.Println(tag[id].topik)
}
