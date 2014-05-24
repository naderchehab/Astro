/** @jsx React.DOM */

var Movie = React.createClass({
    render: function() {
        return (
            <tr>
                <td>{this.props.title}</td>
                <td className="har">{this.props.rank}</td>
                <td>{this.props.year}</td>
                <td>{this.props.rating}</td>
                <td>{this.props.reviews}</td>
            </tr>
            );
    }
});

var MovieList = React.createClass({
    getInitialState: function() {
        return {data: []};
    },
    componentWillMount: function() {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            success: function(data) {
                this.setState({data: data});
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
   render: function() {
       var movies = this.state.data.map(function (movie) {
           return <Movie title={movie.title} rank={movie.rank} year={movie.year} rating={movie.rating} reviews={movie.reviews} />;
       });
       return (
           <tbody>{movies}</tbody>
           );
   }
});

React.renderComponent(
    <table id="movie-table" className="tablesaw" data-mode="stack" data-sortable>
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
        <MovieList url="/public/js/movies.json"/>
    </table>,
    document.getElementById('table-container')
);

$(document).trigger("enhance.tablesaw");