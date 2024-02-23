package main

import (
    _"fmt"
)

type APIServer struct {
    listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
    return &APIServer {
        listenAddr: listenAddr,
    }
}
