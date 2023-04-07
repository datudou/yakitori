'use client';
import React, { useState } from 'react';
import * as echarts from 'echarts';
import ReactECharts from 'echarts-for-react';


function parseGameLog(gameLog: any, quarter: number) {
  const mins = gameLog.mins
  const length = mins[quarter]?.length ?? 0
  if (length === 0) return []

  let start = 0
  const durations = []
  let duration = []

  if ('off' in mins[quarter][0]) {
    start = 720
    duration.push(start)
  } else {
    start = mins[quarter][0]['on']
  }

  for (let i = 0; i < length; i++) {
    if ('off' in mins[quarter][i]) {
      duration.push(mins[quarter][i]['off'])
      durations.push([...duration])
      duration = []
    } else {
      if (i === length - 1) {
        duration.push(mins[quarter][i]['on'], 0)
        durations.push([...duration])
      }
      start = mins[quarter][i]['on']
      duration = [start]
    }
  }
  return durations
}

function renderItem(params, api) {
  let categoryIndex = api.value(0);
  let start = api.coord([12 - api.value(1), categoryIndex]);
  let end = api.coord([12 - api.value(2), categoryIndex]);

  let height = api.size([0, 1])[1] * 0.6;
  let rectShape = echarts.graphic.clipRectByRect(
    {
      x: start[0],
      y: start[1] - height / 2,
      width: end[0] - start[0],
      height: height
    },
    {
      x: params.coordSys.x,
      y: params.coordSys.y,
      width: params.coordSys.width,
      height: params.coordSys.height
    }
  );
  return (
    rectShape && {
      type: 'rect',
      transition: ['shape'],
      shape: rectShape,
      style: api.style()
    }
  );
}

const secondsToMinSecPadded = time => {
  const minutes = "0" + Math.floor(time / 60);
  const seconds = "0" + (time - minutes * 60);
  return minutes.substr(-2) + ":" + seconds.substr(-2);
};


const StackedRangeBar = ({ isHome, gameLogs, quarter, isShowYAxis }) => {
  let cat = []
  let data: { name: string; value: any[]; itemStyle: { normal: { color: string; }; }; }[] = []
  let startTime = 12
  let types = {}
  if (isHome) {
    types = { name: 'ON', color: '#7b9ce1' };
  } else {
    types = { name: 'ON', color: '#fc8452' };
  }
  gameLogs.forEach((gameLog, index) => {
    cat.push(gameLog.player_name)
    let durations = parseGameLog(gameLog, quater);
    if (durations !== null) {
      durations.forEach((duration) => {
        let start = duration[0] / 60
        let end = duration[1] / 60
        data.push({
          name: types.name,
          value: [index, start, end, start - end],
          itemStyle: {
            normal: {
              color: types.color
            },
          }
        });
      });
    }
  });
  const option = {
    tooltip: {
      formatter: function (params:any) {
        console.info(params)
        const time = secondsToMinSecPadded(params.value[3] * 60)
        return params.marker + params.name + ': ' + time + 'm';
      }
    },
    title: {
      text: isHome ? quater + 'th Quarter' : '',
      left: 'center'
    },
    dataZoom: [
      {
        disabled: true,
        type: 'slider',
        filterMode: 'weakFilter',
        showDataShadow: false,
        top: 400,
        labelFormatter: ''
      },
    ],
    grid: {
      left: quater===1 ? 100 : 10, 
      height: 300
    },
    xAxis: {
      min: 0,
      max: 12,
      scale: false,
      axisLabel: {
        formatter: function (val) {
          return Math.min(12, startTime - val) + 'm';
        }
      }
    },
    yAxis: {
      data: cat,
      show: isShowYAxis,
      axisLabel: {
        overflow: "truncate"
      },
      nameTextStyle:{
      }
    },
    series: [
      {
        type: 'custom',
        renderItem: renderItem,
        itemStyle: {
          opacity: 0.8
        },
        encode: {
          x: [1, 2],
          y: 0
        },
        data: data
      }
    ]
  };

  return (
    <>
      <ReactECharts
        option={option}
        style={{ height: 400 }}
      />
      <br />
      <div>
      </div>
    </>
  );

};

export default StackedRangeBar;