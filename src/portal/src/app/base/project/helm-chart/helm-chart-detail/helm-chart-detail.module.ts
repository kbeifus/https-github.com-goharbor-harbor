// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
import { NgModule } from '@angular/core';
import { LabelFilterComponent } from './label-filter/label-filter.component';
import { LabelMarkerComponent } from './label-marker/label-marker.component';
import { ListChartVersionsComponent } from './list-chart-versions/list-chart-versions.component';
import { ChartVersionComponent } from './list-chart-versions/helm-chart-versions-detail/helm-chart-version.component';
import { ChartDetailDependencyComponent } from './chart-detail/chart-detail-dependency.component';
import { ChartDetailSummaryComponent } from './chart-detail/chart-detail-summary.component';
import { ChartDetailValueComponent } from './chart-detail/chart-detail-value.component';
import { ChartDetailComponent } from './chart-detail/chart-detail.component';
import { HelmChartDetailComponent } from './chart-detail.component';
import { SharedModule } from '../../../../shared/shared.module';
import { RouterModule, Routes } from "@angular/router";

const routes: Routes = [
    {
        path: ':chart/versions',
        component: ListChartVersionsComponent
    },
    {
        path: ':chart/versions/:version',
        component: HelmChartDetailComponent,
    },
];
@NgModule({
    imports: [
        SharedModule,
        RouterModule.forChild(routes)
    ],
    declarations: [
        LabelFilterComponent,
        LabelMarkerComponent,
        ListChartVersionsComponent,
        ChartVersionComponent,
        ChartDetailDependencyComponent,
        ChartDetailSummaryComponent,
        ChartDetailValueComponent,
        ChartDetailComponent,
        HelmChartDetailComponent,
    ],
})
export class HelmChartListModule { }
