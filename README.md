# go-sedex-envelop
Version 28.12.2023

This go packages provides definitions for the sedex envelop standard eCH-0090.

# Misc
"Der Bund stellt im Rahmen der Registerharmonisierung eine Plattform für den sicheren Datenaustausch zur Verfügung. Diese Plattform, sedex (steht für: secure data exchange), 
ermöglicht einen sicheren Datenaustausch zwischen den Personenregistern des Bundes und
den kantonalen und kommunalen Einwohnerregistern und die Datenlieferung an das Bundesamt für Statistik.
Unter Federführung des Bundesamts für Statistik wird sedex am 15. Januar 2008 in Betrieb
genommen werden.
Systeme, die an sedex partizipieren sollen, müssen das hier beschriebene Format für Versandumschläge verwenden."[6]

# Receipt Codes
| Code    |   Remark |
|----------|:-------------:|
100 |   Message correctly transmitted
200 |   Invalid envelope syntax
201 |   Duplicate message ID
202 |   No payload found
203 |   Message too old to send
204 |   Message expired
300 |   Unknown sender ID
301 |   Unknown recipient ID
302 |   Unknown physical sender ID
303 |   Invalid message type
304 |   Invalid message class
310 |   Not allowed to send
311 |   Not allowed to receive
312 |   User certificate not valid
313 |   Other recipients are not allowed to receive
330 |   Message size exceeds limit
400 |   Network error
401 |   OSCI hub not reachable
402 |   Folder not reachable
403 |   Logging service not reachable
404 |   Authorization service not reachable
500 |   Internal error
501 |   Error during receiving
601 |   Message successfully sent
701 |   Message expires soon

[7]

# Sources
- [1] bfs, sedex, web, https://www.bfs.admin.ch/bfs/de/home/register/personenregister/sedex.html, 28.12.2023
- [2] sedex, sedex directories, web, https://www.sedex.swiss, 28.12.2023
- [3] eCH, sedex envelop definition, web, http://www.ech.ch/de/ech/ech-0090/2.0, 28.12.2023
- [4] eCH., sedex schema v1, web, https://www.ech.ch/de/xmlns/eCH-0090/1/eCH-0090-1-0.xsd, 28.12.2023
- [5] eCH., sedex schema v2, web,http://www.ech.ch/de/xmlns/eCH-0090/2/eCH-0090-2-0.xsd, 28.12.2023
- [6] Meldewesen, Michel Gentile, Igor Metz, Andres Scheidegger, eCH-0090 – Sedex Umschlag, Web http://www.ech.ch/sites/default/files/dosvers/hauptdokument/STAN_d_DEF_2019-09-19_eCH-0090_V2.0_Sedex%20Umschlag.pdf
- [7] sedex, User documentation for the sedex Client, Web, https://docs.sedex.admin.ch/sedex-client/appendix/receipt-versions/, 28.12.2023
