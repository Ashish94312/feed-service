package data

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomTimestamp() time.Time {
	return time.Now().Add(-time.Duration(rand.Intn(1000000)) * time.Minute)
}

type Post struct {
	ID        string
	AuthorID  string
	Content   string
	Timestamp time.Time
}

var Users = []string{"u1", "u2", "u3", "u4", "u5"}

var Followers = map[string][]string{
	"u1": {"u2", "u3"},
	"u2": {"u3"},
	"u3": {"u1"},
	"u4": {"u5"},
	"u5": {"u1", "u2", "u3", "u4"},
}

// 50 posts for each user
var Posts = map[string][]Post{
	"u2": {
		{ID: "p1", AuthorID: "u2", Content: "Hello from u2", Timestamp: randomTimestamp()},
		{ID: "p2", AuthorID: "u2", Content: "Hello from u2", Timestamp: randomTimestamp()},
		{ID: "p3", AuthorID: "u2", Content: "Hello from u2", Timestamp: randomTimestamp()},
		{ID: "p4", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p5", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p6", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p7", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p8", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p9", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p10", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p11", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p12", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p13", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p14", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p15", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p16", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p17", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p18", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p19", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p20", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p21", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p22", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p23", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p24", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p25", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p26", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p27", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p28", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p29", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p30", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
		{ID: "p31", AuthorID: "u2", Content: "Hello from u2", Timestamp: time.Now().Add(-10 * time.Minute)},
	},

	"u1": {
		{ID: "p32", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p33", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p34", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p35", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p36", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p37", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p38", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p39", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p40", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p41", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p42", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p43", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p44", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p45", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p46", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p47", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p48", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p49", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p50", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p51", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p52", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p53", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p54", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p55", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p56", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p57", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p58", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p59", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p60", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p61", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p62", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p63", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p64", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p65", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
		{ID: "p66", AuthorID: "u1", Content: "Hello from u1", Timestamp: randomTimestamp()},
	},

	"u3": {
		{ID: "p67", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p68", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p69", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p70", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p71", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p72", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p73", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p74", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p75", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p76", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p77", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p78", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p79", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p80", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p81", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p82", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p83", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p84", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p85", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p86", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p87", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p88", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p89", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p90", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p91", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p92", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p93", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p94", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p95", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p96", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p97", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p98", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p99", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p100", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p101", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p102", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p103", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p104", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p105", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p106", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p107", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p108", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p109", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p110", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p111", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p112", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p113", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p114", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p115", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p116", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p117", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p118", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p119", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p120", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p121", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p122", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p123", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p124", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p125", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p126", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p127", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p128", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p129", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p130", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p131", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p132", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p133", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p134", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p135", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p136", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p137", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p138", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p139", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p140", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p141", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
		{ID: "p142", AuthorID: "u3", Content: "Hello from u3", Timestamp: randomTimestamp()},
	},

	"u4": {
		{ID: "p143", AuthorID: "u4", Content: "Hello from u4", Timestamp: randomTimestamp()},
		{ID: "p144", AuthorID: "u4", Content: "Hello from u4", Timestamp: randomTimestamp()},
		{ID: "p145", AuthorID: "u4", Content: "Hello from u4", Timestamp: randomTimestamp()},
	},

	"u5": {
		{ID: "p146", AuthorID: "u5", Content: "Hello from u5", Timestamp: randomTimestamp()},
		{ID: "p147", AuthorID: "u5", Content: "Hello from u5", Timestamp: randomTimestamp()},
		{ID: "p148", AuthorID: "u5", Content: "Hello from u5", Timestamp: randomTimestamp()},
	},
}
