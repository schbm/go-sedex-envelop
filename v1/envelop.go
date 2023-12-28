// Code generated by xgen. DO NOT EDIT.

package v1

import (
	"encoding/xml"
)

// MessageIdType is Datentyp der Meldungs ID. Zeichenkette mit max. 36 Zeichen, die Ziffern, Buchstaben
//         oder Bindestriche enthalten kann. Die Zeichenkette ist lang genug, um eine UUID
//         (vgl. RFC 4122), einen 64 Bit Integer oder eine Art von Schlüssel darzustellen.
//         Beispiele: f81d4fae-7dec-11d0-a765-00a0c91e6bf6, 7F454C4601020100
type MessageIdType string

// ParticipantIdType is Datentyp der sedex Teilnehmer ID.
type ParticipantIdType string

// MessageTypeType is Typ des Meldungstyp.
//         Der Meldungstyp definiert die Funktion eines Datenpakets.
//         Der Wertebereich ist in Nummerierungsbereiche unterteilt.
//         Der Bereich 0000000 -0099999 liegt in der Hoheit von eCH und BFS,
//         die Bereiche nn00000 - nn99999 in der Hoheit des jeweiligen Kantons,
//         wo nn die BFS Nummer des Kantons gemäss historisiertem Gemeindeverzeichnis
//         bezeichnet. Der Meldungstyp definiert
//         zusammen mit der Meldungsklasse (messageClass) implizit, welcher Art
//         (Dateityp bzw. XML Schema) die Nutzdaten der Meldung sind. Der Meldungstyp
//         ist zusammen mit senderId und recipientId für das Routing der Meldung relevant.
type MessageTypeType int

// NameValuePairType is Datentyp der die Übergabe von Daten für Testzwecke.
type NameValuePairType struct {
	XMLName xml.Name `xml:"nameValuePairType"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

// Loopback ...
type Loopback struct {
	XMLName       xml.Name `xml:"loopback"`
	AuthoriseAttr bool     `xml:"authorise,attr"`
}

// EnvelopeType is Versanddatum.Datum (Zeitstempel), an dem die sendende Anwendung die Meldung dem
//         Adapter in den Ausgangsorder gelegt hat.
type EnvelopeType struct {
	XMLName            xml.Name             `xml:"envelopeType"`
	VersionAttr        string               `xml:"version,attr"`
	MessageId          string               `xml:"messageId"`
	MessageType        int                  `xml:"messageType"`
	MessageClass       int                  `xml:"messageClass"`
	ReferenceMessageId string               `xml:"referenceMessageId"`
	SenderId           string               `xml:"senderId"`
	RecipientId        []string             `xml:"recipientId"`
	EventDate          string               `xml:"eventDate"`
	MessageDate        string               `xml:"messageDate"`
	Loopback           *Loopback            `xml:"loopback"`
	TestData           []*NameValuePairType `xml:"testData"`
}

// String255Type is String mit maximaler Länge 255.
type String255Type string

// StatusCodeType is Internal error.Der Adapter kann die Daten nicht senden, weil ein interner
//             Fehler aufgetreten ist. Weitere Informationen stehen im Element
//             statusInfo angefügt.
//             Details zu dem Fehler sind dem Log des Adapters zu entnehmen.
type StatusCodeType int

// ReceiptType is Empfänger der Meldung, auf die sich die Quittung bezieht.
type ReceiptType struct {
	XMLName      xml.Name `xml:"receiptType"`
	VersionAttr  string   `xml:"version,attr"`
	EventDate    string   `xml:"eventDate"`
	StatusCode   int      `xml:"statusCode"`
	StatusInfo   string   `xml:"statusInfo"`
	MessageId    string   `xml:"messageId"`
	MessageType  int      `xml:"messageType"`
	MessageClass int      `xml:"messageClass"`
	SenderId     string   `xml:"senderId"`
	RecipientId  string   `xml:"recipientId"`
}

// Envelope is Umschlag einer sedex Meldung.
type Envelope *EnvelopeType

// Receipt is Versandquittung einer sedex Meldung. Versandquittungen werden vom sedex Adapter
//         ausgestellt. Sie bestätigen, das eine Meldung auf der Gegenseite angekommen ist,
//         oder ob die Meldung nicht ausgeliefert werden konnte. Sie bestätigen aber in
//         keinem Fall, dass die Gegenseite die Meldung verarbeitet hat.
type Receipt *ReceiptType