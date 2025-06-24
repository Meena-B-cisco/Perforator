import React from 'react';

import type { SelectFilter } from 'src/components/Select/Select';
import { uiFactory } from 'src/factory';

import { fetchServices } from './fetchServices';

import './ServiceInput.scss';


export interface ServiceInputProps {
    service?: string;
    onUpdate: (service: string | undefined) => void;
}

export const ServiceInput: React.FC<ServiceInputProps> = props => {
    React.useEffect(() => {
        props.onUpdate(props.service);
    }, []);

    const listValues = React.useCallback(async (filter: SelectFilter) => (
        filter.value
            ? await fetchServices(filter.value, {
                offset: filter.offset,
                limit: filter.limit,
            }) || []
            : []
    ), []);

    return uiFactory().renderSelect({
        value: props.service,
        placeholder: 'Service regexp',
        onUpdate: props.onUpdate,
        listValues,
    });
};


