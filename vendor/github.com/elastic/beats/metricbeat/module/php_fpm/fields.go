// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package php_fpm

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "php_fpm", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzMWEtv4zgSvudXFHKJjXWMmasXWGAxO7PJYWaNdPrUaMg0VbKI8KEmS3bcv35RFCXLspxHx7toHRod0Sx+9dVXD+oWnnC/gKqssqIyVwCkSOMCrpd3y+yP5Z/XVwA5BulVRcrZBfzjCgBgjSS+fL2K/1/eLW//WP4JAf0WPQQSVAcwSF7JANJpjZIwh8I70/54fgUQSucpk84WarOAQuiAVwAeNYqAi3gE/wqJlN2EBXy5DkFfz+C6JKquv14BFAp1HhYRxS1YYbDvCT+0r3ABG+/qKr0ZcYafVdq3AuksCWUDUImdF1QKgh16BLfm1YE7yenOWiU2CFJoPU+v+kiP0Dqnu5djcF+A3MB2Tr8TM/+mxV15JzGEiGPeszyMQvsM/ej7wv8eLbT+POF+53w+WHvBK34eS4wWwRUR8THCn4y/I5uX4DLZzYywYoP+crR+IkFKziDfW2GUBOfB2RyNsPl8FIl01qJkc2EUxZDsN2D4rTMZ8wYhVChVoWT8UwVSMswHu8bIOoAUUmJFOOSixaid3ZwsvQIyKbA2a/SsQWWlM8puwOO3GgMlffTFkApgKUIH6O8jdnclWhA9YkEdNoAi/tOjcdu+4LQKhBa+1VjjkJuWhbh4eQ5k7T1a6nHR00TDQim2CGtEC8oqUoIwn8G6JrCORqzukTqH53DPCa4CbIWukZ23zt5+R++YC9pXigvpHgyK9jih9YhVpklshdJirbENRkokDCB854rew7oO+xkIm/M2j3HVuhGrPQOdbXKN+XikxWdqRTGHB6GCGmGaS42ZG/GcyVLp3KNdwaTybqtyjBhawFJYKIXNNYKiKTfJWudQoh4mGT9PiFVDXgqOdrs5PPKLyrsKPe2hcFq7XThIqRCSIo8jBlspN3ILsFUCBAQnn5Bg8vjbkgtGoTTCWgTMpy2FdQBlS/SKhlWCn+AOdVyWwgtJ6Js859eN+ZOUb2XNpDV4sqjwiwvciGdlatMTeJfjykbQ8VxeqNDm4+Ht50RQViIwj1wLAgnPQj/jXt+1TKO9uHtBfe/YT5Ec+tNHP94GujS4VBNYpsnjQh1A5frywjgIoknOTg+rnVA8lXIlSMyskicTNcf5mbp3KD9sj/eLOCpOU842FTClfCG05pFhx2WVylPvUklRNguV8JgljKuYkm21GSzFE0LTgsZtthNhmjy4Hq/SsLA6J2EhSW3/H/S3rSiBPDAIt83opmzMZAb962rE7CQ4wC1acNyCi5obC0snnTCLVjyGWhPsVAwAxxI8ihxWv6ym5yggR+K0KV1w/GCQf0tMH7LxpZr5P4rKabUcgnpf/ev3xMyjkOWFZ5i/OqCkDIYmxq3MtTKKIsw4vqTzZ+fGtsoAeYUhjgDsFxjHc0XCD5Od808BnNXceIdzOz+VgZuUTzcxU2/a+ftmOl56g3a7rO1Io+V3hJe33K6OeQHRtj3AZ5R1nEx5JbKDzxIxH4nMKm3KGKZ2m4y3uJpWZ5xh0rIokAu58levSkhn8zfJ7xgOQx5FkwsaLhTOG0ELwMrJMmuOfD/3bLgZnZjgs1AH3fcjN91W8B+96v7YzbZzRB2T1bihLOHwlvsiiUzh8v5f3beBE3p68R1E8Pzd+dUTmwZ/fCZM7nONM3iorVV2MwMkOR0HMia0MzJ7VWSvYj0WWB9wT2hncZ7m5wcCZc8k6NtBjRa/iyDqXeYHYHgeeAFLltde8AmXwtTa4xHTKOldS9bkVzBK67QkEonT7j6VRqK2dk/MHtzOQo5FvIs7OybH1guDVLqxlPyhBGkxNFZh8u/fH2ew/M+nxyYtYDKO+SWAtVcj6Aif6cegfX64h52isr3W+T0E8jxIvgMc11C0xLe1DZWXUkCyCo3VNr5dXONMEaEzo+8isw4DLB+IMduCyfJumf3z8+Nd9vnT7w+MxcMtqCLOygFpCpPC+bfCa068SJiNUDbZi9/fMf8gNi0CtcNXJqt6BOXJyPIGlM8gjKstcZQNGuf3zVc+EQ6ZLJ0NtWEPmm+CQu/EPsAv7Eu/XMUvZsTF4T7eZGKXWqMUdcDWuBRa1lq0XxtzZ7G7BHYH9m5VXAIJvVFWEOavEtOc8rMlwn8DAAD//5H9jcU="
}
