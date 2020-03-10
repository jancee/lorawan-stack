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

import PropTypes from '../../../lib/prop-types'

import Field from '../../../components/form/field'
import DurationInput from './schedule-delay-input'

const ScheduleDelayField = function({
  name,
  title,
  required,
  horizontal,
  autoFocus,
  disabled,
  error,
  description,
}) {
  return (
    <Field
      name={name}
      title={title}
      required={required}
      autoFocus={autoFocus}
      component={DurationInput}
      disabled={disabled}
      description={description}
    />
  )
}
ScheduleDelayField.propTypes = {
  autoFocus: PropTypes.bool,
  description: PropTypes.message.isRequired,
  disabled: PropTypes.bool,
  error: PropTypes.error,
  horizontal: PropTypes.bool,
  name: PropTypes.string.isRequired,
  required: PropTypes.bool,
  title: PropTypes.message.isRequired,
}

ScheduleDelayField.defaultProps = {
  required: false,
  horizontal: false,
  autoFocus: false,
  disabled: false,
  error: undefined,
}

export default ScheduleDelayField
