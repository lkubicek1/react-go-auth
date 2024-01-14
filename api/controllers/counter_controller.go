package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"sync/atomic"
)

// The counter is a global variable for simplicity
var counter int64

// UserCounters stores the counters for each user
var UserCounters = struct {
	sync.RWMutex
	counters map[string]int64
}{counters: make(map[string]int64)}

func GetCurrentCounter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"counter": atomic.LoadInt64(&counter)})
}

func IncrementCounter(c *gin.Context) {
	newCounterValue := atomic.AddInt64(&counter, 1)
	c.JSON(http.StatusOK, gin.H{"counter": newCounterValue})
}

func ResetCounter(c *gin.Context) {
	atomic.StoreInt64(&counter, 0)
	c.JSON(http.StatusOK, gin.H{
		"message": "Counter reset successfully",
		"counter": atomic.LoadInt64(&counter),
	})
}

func GetUserCounter(c *gin.Context) {
	// Extract user ID from context, set by AuthMiddleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID required"})
		return
	}

	UserCounters.RLock()
	counter, exists := UserCounters.counters[userID.(string)]
	UserCounters.RUnlock()

	if !exists {
		c.JSON(http.StatusOK, gin.H{"counter": 0})
		return
	}

	c.JSON(http.StatusOK, gin.H{"counter": counter, "userID": userID})
}

func IncrementUserCounter(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID required"})
		return
	}

	UserCounters.Lock()
	UserCounters.counters[userID.(string)]++
	counter := UserCounters.counters[userID.(string)]
	UserCounters.Unlock()

	c.JSON(http.StatusOK, gin.H{"counter": counter, "userID": userID})
}

func ResetUserCounter(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID required"})
		return
	}

	UserCounters.Lock()
	UserCounters.counters[userID.(string)] = 0
	UserCounters.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Counter reset successfully", "counter": 0})
}
