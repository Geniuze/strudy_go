package library

import "errors"

type MusicEntry struct {
    Id string
    Name string
    Artist string
    Source string
    Type string
}

type MusicManager struct {
    musics []MusicEntry
}

func NewMusicManager() *MusicManager {
    return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
    return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
    if index < 0 || index >= len(m.musics) {
        return nil, errors.New("index out of range.")
    }
    return &m.musics[index], nil
}
func (m *MusicManager) Find(name string) (music *MusicEntry, index int, err error) {
    if len(m.musics) == 0 {
        return nil,0, errors.New("have no music")
    }

    for i,m := range m.musics {
        if m.Name == name {
            return &m,i,nil
        }
    }
    return nil,0, nil
}

func (m *MusicManager) Add(music *MusicEntry) {
    m.musics = append(m.musics, *music)
}
func (m *MusicManager) Remove (index int) *MusicEntry {
    if index < 0 || index >= len(m.musics) {
        return nil
    }

    removedMusic := &m.musics[index]

    if index < len(m.musics) - 1 {
        // 中间元素
        m.musics = append(m.musics[:index], m.musics[index+1:]...)
    } else if index == 0 {
        // 删除仅有的一个元素
        m.musics = make([]MusicEntry, 0)
    } else {
        // 删除最后一个元素
        m.musics = m.musics[:index]
    }

    return removedMusic
}
func (m *MusicManager) RemoveByName (name string) *MusicEntry {
    mm,index,_ := m.Find(name)
    if mm != nil {
        return m.Remove(index)
    }
    return nil
}
