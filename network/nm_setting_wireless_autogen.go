// This file is automatically generated, please don't edit manully.
package main

// Get key type
func getSettingWirelessKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_WIRELESS_SSID:
		t = ktypeArrayByte
	case NM_SETTING_WIRELESS_MODE:
		t = ktypeString
	case NM_SETTING_WIRELESS_BAND:
		t = ktypeString
	case NM_SETTING_WIRELESS_CHANNEL:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_BSSID:
		t = ktypeArrayByte
	case NM_SETTING_WIRELESS_RATE:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_TX_POWER:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		t = ktypeArrayByte
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		t = ktypeArrayByte
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_MTU:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SEC:
		t = ktypeString
	case NM_SETTING_WIRELESS_HIDDEN:
		t = ktypeBoolean
	}
	return
}

// Get JSON value generally
func generalGetSettingWirelessKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		LOGGER.Error("generalGetSettingWirelessKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SSID:
		value = getSettingWirelessSsidJSON(data)
	case NM_SETTING_WIRELESS_MODE:
		value = getSettingWirelessModeJSON(data)
	case NM_SETTING_WIRELESS_BAND:
		value = getSettingWirelessBandJSON(data)
	case NM_SETTING_WIRELESS_CHANNEL:
		value = getSettingWirelessChannelJSON(data)
	case NM_SETTING_WIRELESS_BSSID:
		value = getSettingWirelessBssidJSON(data)
	case NM_SETTING_WIRELESS_RATE:
		value = getSettingWirelessRateJSON(data)
	case NM_SETTING_WIRELESS_TX_POWER:
		value = getSettingWirelessTxPowerJSON(data)
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		value = getSettingWirelessMacAddressJSON(data)
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		value = getSettingWirelessClonedMacAddressJSON(data)
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		value = getSettingWirelessMacAddressBlacklistJSON(data)
	case NM_SETTING_WIRELESS_MTU:
		value = getSettingWirelessMtuJSON(data)
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		value = getSettingWirelessSeenBssidsJSON(data)
	case NM_SETTING_WIRELESS_SEC:
		value = getSettingWirelessSecJSON(data)
	case NM_SETTING_WIRELESS_HIDDEN:
		value = getSettingWirelessHiddenJSON(data)
	}
	return
}

// Getter
func getSettingWirelessSsid(data _ConnectionData) (value []byte) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID).([]byte)
	return
}
func getSettingWirelessMode(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE).(string)
	return
}
func getSettingWirelessBand(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND).(string)
	return
}
func getSettingWirelessChannel(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL).(uint32)
	return
}
func getSettingWirelessBssid(data _ConnectionData) (value []byte) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID).([]byte)
	return
}
func getSettingWirelessRate(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE).(uint32)
	return
}
func getSettingWirelessTxPower(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER).(uint32)
	return
}
func getSettingWirelessMacAddress(data _ConnectionData) (value []byte) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS).([]byte)
	return
}
func getSettingWirelessClonedMacAddress(data _ConnectionData) (value []byte) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS).([]byte)
	return
}
func getSettingWirelessMacAddressBlacklist(data _ConnectionData) (value []string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST).([]string)
	return
}
func getSettingWirelessMtu(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU).(uint32)
	return
}
func getSettingWirelessSeenBssids(data _ConnectionData) (value []string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS).([]string)
	return
}
func getSettingWirelessSec(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC).(string)
	return
}
func getSettingWirelessHidden(data _ConnectionData) (value bool) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN).(bool)
	return
}

// Setter
func setSettingWirelessSsid(data _ConnectionData, value []byte) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, value)
}
func setSettingWirelessMode(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, value)
}
func setSettingWirelessBand(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, value)
}
func setSettingWirelessChannel(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, value)
}
func setSettingWirelessBssid(data _ConnectionData, value []byte) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, value)
}
func setSettingWirelessRate(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, value)
}
func setSettingWirelessTxPower(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, value)
}
func setSettingWirelessMacAddress(data _ConnectionData, value []byte) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, value)
}
func setSettingWirelessClonedMacAddress(data _ConnectionData, value []byte) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, value)
}
func setSettingWirelessMacAddressBlacklist(data _ConnectionData, value []string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, value)
}
func setSettingWirelessMtu(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, value)
}
func setSettingWirelessSeenBssids(data _ConnectionData, value []string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, value)
}
func setSettingWirelessSec(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, value)
}
func setSettingWirelessHidden(data _ConnectionData, value bool) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, value)
}

// JSON Getter
func getSettingWirelessSsidJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SSID))
	return
}
func getSettingWirelessModeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MODE))
	return
}
func getSettingWirelessBandJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BAND))
	return
}
func getSettingWirelessChannelJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CHANNEL))
	return
}
func getSettingWirelessBssidJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BSSID))
	return
}
func getSettingWirelessRateJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, getSettingWirelessKeyType(NM_SETTING_WIRELESS_RATE))
	return
}
func getSettingWirelessTxPowerJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, getSettingWirelessKeyType(NM_SETTING_WIRELESS_TX_POWER))
	return
}
func getSettingWirelessMacAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS))
	return
}
func getSettingWirelessClonedMacAddressJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS))
	return
}
func getSettingWirelessMacAddressBlacklistJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST))
	return
}
func getSettingWirelessMtuJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MTU))
	return
}
func getSettingWirelessSeenBssidsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEEN_BSSIDS))
	return
}
func getSettingWirelessSecJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEC))
	return
}
func getSettingWirelessHiddenJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, getSettingWirelessKeyType(NM_SETTING_WIRELESS_HIDDEN))
	return
}

// JSON Setter
func setSettingWirelessSsidJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SSID))
}
func setSettingWirelessModeJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MODE))
}
func setSettingWirelessBandJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BAND))
}
func setSettingWirelessChannelJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CHANNEL))
}
func setSettingWirelessBssidJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BSSID))
}
func setSettingWirelessRateJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_RATE))
}
func setSettingWirelessTxPowerJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_TX_POWER))
}
func setSettingWirelessMacAddressJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS))
}
func setSettingWirelessClonedMacAddressJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS))
}
func setSettingWirelessMacAddressBlacklistJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST))
}
func setSettingWirelessMtuJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MTU))
}
func setSettingWirelessSeenBssidsJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEEN_BSSIDS))
}
func setSettingWirelessSecJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEC))
}
func setSettingWirelessHiddenJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_HIDDEN))
}

// Remover
func removeSettingWirelessSsid(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID)
}
func removeSettingWirelessMode(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE)
}
func removeSettingWirelessBand(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND)
}
func removeSettingWirelessChannel(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL)
}
func removeSettingWirelessBssid(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID)
}
func removeSettingWirelessRate(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE)
}
func removeSettingWirelessTxPower(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER)
}
func removeSettingWirelessMacAddress(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS)
}
func removeSettingWirelessClonedMacAddress(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS)
}
func removeSettingWirelessMacAddressBlacklist(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST)
}
func removeSettingWirelessMtu(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU)
}
func removeSettingWirelessSeenBssids(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS)
}
func removeSettingWirelessSec(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC)
}
func removeSettingWirelessHidden(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN)
}