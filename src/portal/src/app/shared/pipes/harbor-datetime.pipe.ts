import { Pipe, PipeTransform } from '@angular/core';
import { DatePipe } from "@angular/common";
import { getDatetimeRendering } from "../units/shared.utils";
import { DEFAULT_LANG_LOCALSTORAGE_KEY, DeFaultLang, DatetimeRendering } from "../entities/shared.const";

const baseTimeLine: Date = new Date('1970-1-1');

const formatTransformers: Record<DatetimeRendering, (format: string) => string> = {
    "iso-8601": asISO8601,
    "locale-default": format => format,
} as const;

@Pipe({
    name: 'harborDatetime',
    pure: false
})
export class HarborDatetimePipe implements PipeTransform {

    transform(value: any, format?: string): string {
        let lang: string = DeFaultLang;
        if (localStorage && localStorage.getItem(DEFAULT_LANG_LOCALSTORAGE_KEY)) {
            lang = localStorage.getItem(DEFAULT_LANG_LOCALSTORAGE_KEY);
        }
        if (value && value <= baseTimeLine) {// invalid date
            return '-';
        }
        const formatTransformer = formatTransformers[getDatetimeRendering()];
        // default format medium
        return new DatePipe(lang).transform(value, formatTransformer(format ? format : 'medium'));
    }
}

function asISO8601<Format extends string>(format: Format) {
    switch (format) {
        // https://angular.io/api/common/DatePipe#pre-defined-format-options
        case "short":      return "yyyy-MM-dd, HH:mm";
        case "medium":     return "yyyy-MM-dd, HH:mm:ss";
        case "long":       return "yyyy-MM-dd, HH:mm:ss z";
        case "full":       return "EEEE yyyy-MM-dd, HH:mm:ss zzzz";
        case "shortDate":  return "yyyy-MM-dd";
        case "mediumDate": return "yyyy-MM-dd";
        case "longDate":   return "yyyy-MM-dd z";
        case "fullDate":   return "EEEE yyyy-MM-dd zzzz";
        case "shortTime":  return "HH:mm";
        case "mediumTime": return "HH:mm:ss";
        case "longTime":   return "HH:mm:ss z";
        case "fullTime":   return "HH:mm:ss zzzz";
        default:           return format;
    }
}
