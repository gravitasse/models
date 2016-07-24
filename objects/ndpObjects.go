//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

type NDPGlobal struct {
	baseObj
	// placeholder to create a key
	Vrf    string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "System Vrf", DEFAULT:"default"`
	Enable string `DESCRIPTION: "Enable/Diable Neighbor Discovery Globally", DEFAULT: "true"`
}

type NDPEntryState struct {
	baseObj
	IpAddr         string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Neighbor's IP Address"`
	MacAddr        string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address"`
	Vlan           string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to"`
	Intf           string `DESCRIPTION: "Router Interface to which neighbor is attached to"`
	ExpiryTimeLeft string `DESCRIPTION: "Time left before entry expires in case neighbor departs"`
}
