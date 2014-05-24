/** @jsx React.DOM */
React.renderComponent(
    <table id="t1" className="tablesaw" data-mode="swipe" data-sortable data-sortable-switch data-minimap data-mode-switch>
        <thead>
            <tr>
                <th data-sortable-col data-sortable-default-col data-priority="persist">Movie Title</th>
                <th data-sortable-col data-priority="3">Rank</th>
                <th data-sortable-col data-priority="2">Year</th>
                <th data-sortable-col data-priority="1">
                    <abbr title="Rotten Tomato Rating">Rating</abbr>
                </th>
                <th data-sortable-col data-priority="4">Reviews</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>
                    <a href="http://en.wikipedia.org/wiki/Citizen_Kane" data-rel="external">Citizen Kane</a>
                </td>
                <td class="har">1</td>
                <td>1941</td>
                <td>100%</td>
                <td>74</td>
            </tr>
        </tbody>
    </table>,
    document.getElementById('table-container')
);

$(document).trigger("enhance.tablesaw");