package header

func Render() string {
	return `
        <nav class="blue">
            <div class="nav-wrapper container">
                <a href="#" class="brand-logo">Go</a>
                <ul id="nav-mobile" class="right hide-on-med-and-down">
                    <li><a href="#infoSection">Info</a></li>
                    <li><a href="#footer">Footer</a></li>
                </ul>
            </div>
        </nav>
    `
}
