package asbclient

import (
	"encoding/xml"
	"time"
)

type QueueFeed struct {
	XMLName xml.Name         `xml:"feed"`
	Entry   []QueueFeedEntry `xml:"entry"`
}

type QueueFeedEntry struct {
	XMLName xml.Name              `xml:"entry"`
	Title   string                `xml:"title"`
	ID      string                `xml:"id"`
	Content QueueFeedEntryContent `xml:"content"`
}
type QueueFeedEntryContent struct {
	XMLName          xml.Name `xml:"content"`
	QueueDescription QueueDescription
}

//QueueDescription https://docs.microsoft.com/en-us/dotnet/api/microsoft.servicebus.messaging.queuedescription?view=azure-dotnet
type QueueDescription struct {
	XMLName                             xml.Name          `xml:"QueueDescription"`
	LockDuration                        string            //	<LockDuration>PT1M</LockDuration>
	MaxSizeInMegabytes                  uint              //	<MaxSizeInMegabytes>1024</MaxSizeInMegabytes>
	RequiresDuplicateDetection          bool              //	<RequiresDuplicateDetection>false</RequiresDuplicateDetection>
	RequiresSession                     bool              //	<RequiresSession>false</RequiresSession>
	DefaultMessageTimeToLive            string            //	<DefaultMessageTimeToLive>P10675199DT2H48M5.4775807S</DefaultMessageTimeToLive>
	DeadLetteringOnMessageExpiration    bool              //	<DeadLetteringOnMessageExpiration>false</DeadLetteringOnMessageExpiration>
	DuplicateDetectionHistoryTimeWindow string            //	<DuplicateDetectionHistoryTimeWindow>PT10M</DuplicateDetectionHistoryTimeWindow>
	MaxDeliveryCount                    uint              //	<MaxDeliveryCount>10</MaxDeliveryCount>
	EnableBatchedOperations             bool              //	<EnableBatchedOperations>false</EnableBatchedOperations>
	SizeInBytes                         uint              //	<SizeInBytes>518051</SizeInBytes>
	MessageCount                        uint              //	<MessageCount>511</MessageCount>
	IsAnonymousAccessible               bool              //	<IsAnonymousAccessible>false</IsAnonymousAccessible>
	Status                              string            //	<Status>Active</Status>
	ForwardTo                           string            //	<ForwardTo/>
	CreatedAt                           XmlTime           //	<CreatedAt>2018-02-02T09:36:25.42Z</CreatedAt>
	UpdatedAt                           XmlTime           //	<UpdatedAt>2018-02-02T09:36:25.48Z</UpdatedAt>
	AccessedAt                          XmlTime           //	<AccessedAt>2018-02-19T04:44:31.8691034Z</AccessedAt>
	SupportOrdering                     bool              //	<SupportOrdering>false</SupportOrdering>
	CountDetails                        QueueCountDetails //	<CountDetails xmlns:d2p1="http://schemas.microsoft.com/netservices/2011/06/servicebus">
	//	</CountDetails>
	AutoDeleteOnIdle              string //	<AutoDeleteOnIdle>P10675199DT2H48M5.4775807S</AutoDeleteOnIdle>
	EnablePartitioning            bool   //	<EnablePartitioning>false</EnablePartitioning>
	EntityAvailabilityStatus      string //	<EntityAvailabilityStatus>Available</EntityAvailabilityStatus>
	ForwardDeadLetteredMessagesTo string //	<ForwardDeadLetteredMessagesTo/>
	EnableExpress                 bool   //	<EnableExpress>false</EnableExpress>
	//	</QueueDescription>
}

//QueueCountDetails https://docs.microsoft.com/en-us/dotnet/api/microsoft.servicebus.messaging.queuedescription?view=azure-dotnet
type QueueCountDetails struct {
	//	<CountDetails xmlns:d2p1="http://schemas.microsoft.com/netservices/2011/06/servicebus">
	ActiveMessageCount             uint //		<d2p1:ActiveMessageCount>510</d2p1:ActiveMessageCount>
	DeadLetterMessageCount         uint //		<d2p1:DeadLetterMessageCount>1</d2p1:DeadLetterMessageCount>
	ScheduledMessageCount          uint //		<d2p1:ScheduledMessageCount>0</d2p1:ScheduledMessageCount>
	TransferMessageCount           uint //		<d2p1:TransferMessageCount>0</d2p1:TransferMessageCount>
	TransferDeadLetterMessageCount uint //		<d2p1:TransferDeadLetterMessageCount>0</d2p1:TransferDeadLetterMessageCount>
	//	</CountDetails>

}

type XmlTime struct {
	time.Time
}

func (c *XmlTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*c = XmlTime{parse}
	return nil
}
