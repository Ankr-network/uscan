/*
Copyright © 2022 uscan team

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package share

const (
	HttpAddr = "http_addr"
	HttpPort = "http_port"
	TlsPath  = "tls_path"
	TLS      = "tls"

	RpcUrls  = "rpc_urls"
	WorkChan = "work_chan"

	MdbxPath = "db_path"

	ForkBlockNum = "fork_block_number"

	APPTitle    = "app_title"    //app_title是用户自定义的浏览器标题，比如是Coq的话就显示 Coq Chain Scan
	UnitDisplay = "unit_display" //unit_display是用户指定显示的单位，比如是Eth、Peel、Bnb
	NodeUrl     = "node_url"     //node_url是需要和合约交互的时候使用的节点
)
