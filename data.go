package ml

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"

	"github.com/ornerylawn/linear"
)

func LoadMatrixCSV(filename string) (linear.Matrix, error) {
	f, err := os.Open(f