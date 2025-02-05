/*
 * Copyright 2021 LimeChain Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package timestamp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	nanosInSecond = 1000000000
)

func FromString(timestamp string) (int64, error) {
	var err error
	stringTimestamp := strings.Split(timestamp, ".")

	seconds, err := strconv.ParseInt(stringTimestamp[0], 10, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Could not parse the whole part of a timestamp: [%s] - [%s]", timestamp, err))
	}
	nano, err := strconv.ParseInt(stringTimestamp[1], 10, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Could not parse the decimal part of a timestamp: [%s] - [%s]", timestamp, err))
	}

	return seconds*nanosInSecond + nano, nil
}

func ToString(timestamp int64) string {
	seconds := timestamp / nanosInSecond
	nano := timestamp % nanosInSecond
	return fmt.Sprintf("%d.%d", seconds, nano)
}

// ToNanos - converts timestamp in seconds to timestamp in nanos
func ToNanos(timestampInSec int64) int64 {
	return timestampInSec * nanosInSecond
}
