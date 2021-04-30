/*
 * @Author: yhlyl
 * @Date: 2021-04-27 09:53:32
 * @LastEditTime: 2021-04-27 14:06:48
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /yhgolang/eth/eth/url_test.go
 */

package eth

import "testing"

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://ethereum.org")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "ethereum.org" {
		t.Errorf("expected: %v, got: %v", "ethereum.org", url.Path)
	}

	_, err = parseURL("ethereum.org")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "ethereum.org"}
	if url.String() != "https://ethereum.org" {
		t.Errorf("expected: %v, got: %v", "https://ethereum.org", url.String())
	}

	url = URL{Scheme: "", Path: "ethereum.org"}
	if url.String() != "ethereum.org" {
		t.Errorf("expected: %v, got: %v", "ethereum.org", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "ethereum.org"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://ethereum.org\"" {
		t.Errorf("expected: %v, got: %v", "\"https://ethereum.org\"", string(json))
	}
}
