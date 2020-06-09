package model

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/signalfx/com_signalfx_metrics_protobuf"
)

func TestUnmarshalMarshal(t *testing.T) {
	model := generateDataPointUploadMessage()
	bytes, _ := model.Marshal()
	proto2Model := &com_signalfx_metrics_protobuf.DataPointUploadMessage{}
	if err := proto.Unmarshal(bytes, proto2Model); err != nil {
		t.Fail()
	}
	var lightBytes []byte
	var err error
	if lightBytes, err = proto.Marshal(proto2Model); err != nil {
		t.Fail()
	}
	newModel := &DataPointUploadMessage{}
	if err := newModel.Unmarshal(lightBytes); err != nil {
		t.Fail()
	}
	assert.EqualValues(t, model, newModel)
}

func BenchmarkProto2Model(b *testing.B) {
	uploadMessage := generateDataPointUploadMessage()
	bytes, _ := uploadMessage.Marshal()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		proto2Model := &com_signalfx_metrics_protobuf.DataPointUploadMessage{}
		if err := proto.Unmarshal(bytes, proto2Model); err != nil {
			b.Fail()
		}
		if _, err := proto.Marshal(proto2Model); err != nil {
			b.Fail()
		}
	}
}

func BenchmarkModel(b *testing.B) {
	uploadMessage := generateDataPointUploadMessage()
	bytes, _ := uploadMessage.Marshal()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		model := &DataPointUploadMessage{}
		if err := model.Unmarshal(bytes); err != nil {
			b.Fail()
		}
		if _, err := model.Marshal(); err != nil {
			b.Fail()
		}
	}
}

func generateDataPointUploadMessage() *DataPointUploadMessage {
	ret := &DataPointUploadMessage{}
	ret.Datapoints = append(ret.Datapoints, generateEnumMetrics(10)...)
	ret.Datapoints = append(ret.Datapoints, generateGaugeMetrics(10)...)
	ret.Datapoints = append(ret.Datapoints, generateCounterMetrics(10)...)
	return ret
}

func generateEnumMetrics(numPoints int) []*DataPoint {
	val := "123"
	mType := MetricType_ENUM
	ret := make([]*DataPoint, 0, numPoints)
	for i := 0; i < numPoints; i++ {
		ret = append(ret,
			&DataPoint{
				Source:    "mySource",
				Metric:    "myMetricEnum",
				Timestamp: 100,
				Value: Datum{
					StrValue: &val,
				},
				MetricType: &mType,
				Dimensions: []*Dimension{
					{
						Key:   "myKey1",
						Value: "myValue2",
					},
					{
						Key:   "myKey2",
						Value: "myValue2",
					},
				},
			})
	}
	return ret
}

func generateGaugeMetrics(numPoints int) []*DataPoint {
	val := 12.3
	mType := MetricType_GAUGE
	ret := make([]*DataPoint, 0, numPoints)
	for i := 0; i < numPoints; i++ {
		ret = append(ret,
			&DataPoint{
				Source:    "mySource",
				Metric:    "myMetricGauge",
				Timestamp: 100,
				Value: Datum{
					DoubleValue: &val,
				},
				MetricType: &mType,
				Dimensions: []*Dimension{
					{
						Key:   "myKey1",
						Value: "myValue2",
					},
					{
						Key:   "myKey2",
						Value: "myValue2",
					},
				},
			})
	}
	return ret
}

func generateCounterMetrics(numPoints int) []*DataPoint {
	mType := MetricType_COUNTER
	val := int64(123)
	ret := make([]*DataPoint, 0, numPoints)
	for i := 0; i < numPoints; i++ {
		ret = append(ret, &DataPoint{
			Source:    "mySource",
			Metric:    "myMetricCounter",
			Timestamp: 100,
			Value: Datum{
				IntValue: &val,
			},
			MetricType: &mType,
			Dimensions: []*Dimension{
				{
					Key:   "myKey1",
					Value: "myValue2",
				},
				{
					Key:   "myKey2",
					Value: "myValue2",
				},
			},
		})
	}
	return ret
}
