// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import classnames from 'classnames'
import bind from 'autobind-decorator'
import { defineMessages } from 'react-intl'

import Select from '../../../components/select'
import Input from '../../../components/input'

import PropTypes from '../../../lib/prop-types'

import Notification from '../../../components/notification'
import style from './schedule-delay-input.styl'

const m = defineMessages({
  miliseconds: 'miliseconds',
  seconds: 'seconds',
  minutes: 'minutes',
  hours: 'hours',
  durationWarning:
    'Value of delay too low. The lower bound (130ms) will be used by the Gateway Server.',
})

const timeUnitOptions = [
  { label: m.miliseconds, value: 'ms' },
  { label: m.seconds, value: 's' },
  { label: m.minutes, value: 'm' },
  { label: m.hours, value: 'h' },
]

class ScheduleDelayInput extends React.PureComponent {
  constructor(props) {
    super(props)
    const { value } = this.props
    let initialValue

    if (!value) {
      initialValue = {
        duration: 530,
        unit: 'ms',
      }
    } else {
      initialValue = {
        duration: Number(value.substring(0, value.length - 1)),
        unit: 's',
      }
    }

    this.state = {
      duration: initialValue.duration,
      unit: initialValue.unit,
    }
  }

  @bind
  async handleChange(value) {
    const { onChange } = this.props
    const { unit } = this.state

    this.setState({ duration: value })
    // if duration is empty, propagate empty value to form
    if (value) {
      onChange(`${value}${unit}`)
    } else {
      onChange('')
    }
  }

  @bind
  async handleUnitChange(value) {
    const { onChange } = this.props
    const { duration } = this.state
    await this.setState({ unit: value })
    // if duration is empty, propagate empty value to form
    if (duration) {
      onChange(`${duration}${value}`)
    } else {
      onChange('')
    }
  }

  isValidDuration() {
    const { duration, unit } = this.state
    const durationLowerBound = 130 // set lower bound duration in ms
    switch (unit) {
      case 'ms':
        return duration >= durationLowerBound
      case 's':
        return duration >= durationLowerBound / 1000
      case 'm':
        return duration >= durationLowerBound / 60000
      case 'h':
        return duration >= durationLowerBound / 3600000
    }
  }

  render() {
    const { className, name } = this.props
    const { duration, unit } = this.state
    const selectTimeUnitComponent = (
      <Select
        className={style.select}
        name={`${name}-select`}
        options={timeUnitOptions}
        onChange={this.handleUnitChange}
        value={unit}
      />
    )

    return (
      <React.Fragment>
        <div className={classnames(className, style.container)}>
          <Input
            className={style.number}
            ref={this.inputRef}
            type="number"
            step="any"
            name={name}
            onBlur={this.handleBlur}
            value={duration}
            onChange={this.handleChange}
          />
          {selectTimeUnitComponent}
        </div>
        {!this.isValidDuration() && (
          <Notification className={style.notification} small content={m.durationWarning} warning />
        )}
      </React.Fragment>
    )
  }
}

ScheduleDelayInput.propTypes = {
  className: PropTypes.string,
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func.isRequired,
  value: PropTypes.string,
}

ScheduleDelayInput.defaultProps = {
  className: undefined,
  value: undefined,
}

export default ScheduleDelayInput
